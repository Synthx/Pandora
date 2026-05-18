package login

import (
	"bufio"
	"context"
	"net"
	login_message "pandora/login/message"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	saltLength = 32
)

type Client struct {
	Id   string
	Salt string
	conn *net.TCPConn
	log  *zap.Logger
}

func NewClient(conn *net.TCPConn, logger *zap.Logger) (*Client, error) {
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
		Id:   id,
		Salt: salt,
		conn: conn,
		log:  logger.Named("login_client").With(zap.String("client_id", id)),
	}, nil
}

func (c *Client) listen(ctx context.Context) error {
	reader := bufio.NewReader(c.conn)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			message, err := reader.ReadString('\x00')
			if err != nil {
				return err
			}

			message = strings.TrimSuffix(message, "\n\x00")
			c.log.Debug("Message received", zap.String("raw_message", message))
		}
	}
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		c.log.Error("Failed to close connection", zap.Error(err))
	}

	c.log.Debug("Connection closed")
}

func (c *Client) Send(message login_message.OutboundMessage) error {
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
