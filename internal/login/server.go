package login

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"pandora/internal/login/message"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	port       int
	listener   *net.TCPListener
	clients    map[string]*Client
	dispatcher *MessageDispatcher
	log        *zap.Logger
	mutex      sync.Mutex
	wg         sync.WaitGroup
}

func NewServer(port int, logger *zap.Logger, dispatcher *MessageDispatcher) (*Server, error) {
	addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		return nil, err
	}

	server := &Server{
		port:       port,
		listener:   listener,
		clients:    make(map[string]*Client),
		dispatcher: dispatcher,
		log:        logger.Named("server"),
	}

	return server, nil
}

func (s *Server) Serve(ctx context.Context) error {
	s.log.Info(fmt.Sprintf("Listening on port %d", s.port))

	go func() {
		<-ctx.Done()
		s.log.Info("Shutting down server...")
		_ = s.listener.Close()
	}()

	for {
		conn, err := s.listener.AcceptTCP()
		if err != nil {
			select {
			case <-ctx.Done():
				s.wg.Wait()
				s.log.Info("All connections closed. Shutdown complete")

				return nil
			default:
				s.log.Error("Error accepting connection", zap.Error(err))
				continue
			}
		}

		s.wg.Add(1)

		go func() {
			defer s.wg.Done()
			err := s.handleConnection(ctx, conn)
			if err != nil && !errors.Is(err, io.EOF) {
				s.log.Error("Error handling connection", zap.Error(err))
			}
		}()
	}
}

func (s *Server) addClient(client *Client) {
	s.mutex.Lock()
	s.clients[client.Id] = client
	s.mutex.Unlock()
}

func (s *Server) removeClient(client *Client) {
	s.mutex.Lock()
	delete(s.clients, client.Id)
	s.mutex.Unlock()
}

func (s *Server) handleConnection(ctx context.Context, conn *net.TCPConn) error {
	client, err := NewClient(conn, s.log, s.dispatcher)
	if err != nil {
		return err
	}

	defer func() {
		client.Close()
		s.removeClient(client)
	}()

	s.addClient(client)

	err = client.Send(message.NewHelloConnectMessage(client.Salt))
	if err != nil {
		return err
	}

	return client.listen(ctx)
}
