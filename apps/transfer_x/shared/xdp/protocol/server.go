package xdp

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/dehwyy/acheron/apps/stream_x/shared/xdp/protocol/workerpool"
)

// @TLS 1.3 Example
// cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
//  config := &tls.Config{
//  	Certificates: []tls.Certificate{cert},
//  	MinVersion:   tls.VersionTLS13, // Устанавливаем минимум TLS 1.3
//  }

type ServerParams struct {
	tls *tls.Config
}

type Server struct {
	tcpListener net.Listener
	workerPool  workerpool.WorkerPool
}

func NewXDPServer(params ServerParams) (*Server, error) {
	var listener net.Listener

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{})
	if err != nil {
		return nil, err
	}

	listener = tls.NewListener(listener, params.tls)

	return &Server{
		tcpListener: listener,
		workerPool:  workerpool.NewDefaultWorkerPool(),
	}, nil
}

func (s *Server) Start() {
	ctx := context.Background()
	s.workerPool.StartWorkers(ctx)

	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {

			return
		}

		// ? Should I remove this label?
	selectLoop:
		for {
			select {
			case <-time.After(1 * time.Second):
				// Logger typeshit
			case err = <-s.workerPool.QueueConnection(conn):
				// Logger typeshit
				break selectLoop
			}
		}

	}
}

func (s *Server) Stop() {
	s.tcpListener.Close()
}
