package workerpool

import (
	"context"
	"net"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/connection"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
)

const (
	defaultWorkersCount uint = 8
)

func newWorker(ctx context.Context, r router.Router, connectionChannel <-chan net.Conn) {

	connHandler := connection.NewConnectionHandler(r)

	for {
		select {
		case <-ctx.Done():
			return
		case conn, ok := <-connectionChannel:
			if !ok {
				return
			}

			defer conn.Close() // ! Remove or not?

			err := connHandler.HandleConnection(conn)
			if err != nil {
				// Logger typeshit
			}
		}
	}
}

type DefaultWorkerPool struct {
	connectionChannel chan net.Conn
}

// func NewDefaultWorkerPool() WorkerPool {
// 	return &DefaultWorkerPool{
// 		connectionChannel: nil,
// 	}
// }

func (p *DefaultWorkerPool) StartWorkers(ctx context.Context, r router.Router, workers ...uint) {
	if p.connectionChannel != nil {
		close(p.connectionChannel)
	}

	workersCount := defaultWorkersCount
	if len(workers) > 0 {
		workersCount = workers[0]
	}

	p.connectionChannel = make(chan net.Conn, workersCount)

	for i := uint(0); i < workersCount; i++ {
		go newWorker(ctx, r, p.connectionChannel)
	}
}

func (p *DefaultWorkerPool) Stop() {
	if p.connectionChannel == nil {
		return
	}

	close(p.connectionChannel)
}

func (p *DefaultWorkerPool) QueueConnection(conn net.Conn) <-chan error {
	ch := make(chan error, 1)
	p.connectionChannel <- conn
	ch <- nil
	return ch
}
