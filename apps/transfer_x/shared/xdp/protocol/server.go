package xdp

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/workerpool"
)

// @TLS 1.3 Example
// cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
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

	// TODO: add field for &net.TCPAddr
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{})
	if err != nil {
		return nil, err
	}

	listener = tls.NewListener(listener, params.TLS)

	return &Server{
		tcpListener: listener,
		workerPool:  workerpool.NewDefaultWorkerPool(),
	}, nil
}

func (s *Server) Start() error {

	ctx := context.Background()
	s.workerPool.StartWorkers(ctx)

	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			return err
		}

		// ? Should I remove this label?
	selectLoop:
		for {
			select {
			case <-time.NewTimer(1 * time.Second).C:
				// Logger typeshit
			case err = <-s.workerPool.QueueConnection(conn):
				if err != nil {
					return err
				}
				// Logger typeshit
				break selectLoop
			}
		}
	}
}

func (s *Server) Stop() {
	s.workerPool.Stop()
	s.tcpListener.Close()
}
