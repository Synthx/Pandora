package login

import (
	"bufio"
	"context"
	"net"
	"pandora/internal/login/message"
	"pandora/internal/pkg"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	saltLength = 32
)

type ClientState int

const (
	StateWaitingVersion ClientState = iota
	StateWaitingCredentials
	StateWaitingQueuePosition
	StateWaitingNickname
	StateIdle
)

type Client struct {
	Id         string
	Salt       string
	Version    string
	Token      string
	state      ClientState
	conn       *net.TCPConn
	dispatcher *MessageDispatcher
	log        *zap.Logger
}

func NewClient(conn *net.TCPConn, logger *zap.Logger, dispatcher *MessageDispatcher) (*Client, error) {
	err := conn.SetNoDelay(true)
	if err != nil {
		return nil, err
	}

	err = conn.SetKeepAlive(true)
	if err != nil {
		return nil, err
	}

	err = conn.SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return nil, err
	}

	id := conn.RemoteAddr().String()
	salt, err := RandomSalt(saltLength)
	if err != nil {
		return nil, err
	}

	return &Client{
		Id:         id,
		Salt:       salt,
		state:      StateWaitingVersion,
		conn:       conn,
		dispatcher: dispatcher,
		log:        logger.Named("login_client").With(zap.String("client_id", id)),
	}, nil
}
func (c *Client) listen(ctx context.Context) error {
	reader := bufio.NewReader(c.conn)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			packet, err := reader.ReadString('\x00')
			if err != nil {
				return err
			}

			packet = strings.TrimSuffix(packet, "\n\x00")
			if len(packet) == 0 {
				continue
			}

			c.log.Debug("Message received", zap.String("packet", packet))

			c.handleMessage(ctx, packet)
		}
	}
}

func (c *Client) handleMessage(ctx context.Context, rawMessage string) {
	switch c.state {
	case StateWaitingVersion:
		message, err := message.NewClientVersionMessage(rawMessage)
		if err != nil {
			c.Close()
			return
		}

		c.Version = message.Version
		c.state = StateWaitingCredentials
	case StateWaitingCredentials:
		_, err := message.NewClientCredentialsMessage(rawMessage)
		if err != nil {
			c.log.Error("Failed to parse credentials", zap.Error(err))
		}

		c.Token = "test"
		c.state = StateIdle
	case StateWaitingQueuePosition:
		c.state = StateIdle
	case StateWaitingNickname:
		c.state = StateIdle
	case StateIdle:
		err := c.dispatcher.Dispatch(ctx, c, rawMessage)
		if err != nil {
			c.log.Error("Error while processing message", zap.Error(err))
		}

		return
	}
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		c.log.Error("Failed to close connection", zap.Error(err))
	}

	c.log.Debug("Connection closed")
}

func (c *Client) Send(message pkg.OutboundMessage) error {
	serialized, err := message.Serialize()
	if err != nil {
		return err
	}

	rawMessage := message.GetHeader() + serialized
	_, err = c.conn.Write([]byte(rawMessage + "\x00"))
	if err != nil {
		return err
	}

	c.log.Debug("Message sent", zap.String("raw_message", rawMessage))

	return nil
}
