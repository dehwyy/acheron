package xdp

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/workerpool"
)

// @TLS 1.3 Example
// cert, err := tls.LoadX509KeyPair("cerpem", "key.pem")
//  config := &tls.Config{
//  	Certificates: []tls.Certificate{cert},
//  	MinVersion:   tls.VersionTLS13, // TLS 1.3
//  }

type ServerParams struct {
	TLS *tls.Config
}

type Server struct {
	tcpListener net.Listener
	workerPool  workerpool.WorkerPool
}

func NewXDPServer(params ServerParams) (*Server, error) {
	var listener net.Listener

	// TODO: add fields for &net.TCPAddr
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{})
	if err != nil {
		return nil, err
	}

	listener = tls.NewListener(listener, params.TLS)

	srv := &Server{
		tcpListener: listener,
		workerPool:  workerpool.NewWorkerPool(),
	}

	return srv, nil
}

func (s *Server) Start(r router.Router) error {
	ctx := context.Background()
	s.workerPool.StartWorkers(ctx, r)

	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			return err
		}

		// ? Should I remove this label?
		for {
			select {
			case <-time.NewTimer(1 * time.Second).C:
				log.Logger.Debug().Msgf("Connection not handled yet!")
			case err = <-s.workerPool.QueueConnection(conn):
				if err != nil {
					log.Logger.Error().Msgf("Failed to queue connection: %v", err)
					return err
				}
			}
		}
	}
}

func (s *Server) Stop() {
	s.workerPool.Stop()
	if err := s.tcpListener.Close(); err != nil {
		log.Logger.Error().Msgf("Failed to close listener: %v", err)
	}
}
