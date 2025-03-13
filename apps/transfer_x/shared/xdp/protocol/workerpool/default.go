package workerpool

import (
	"context"
	"net"
)

const (
	defaultWorkersCount uint = 8
)

func newWorker(ctx context.Context, connectionChannel <-chan net.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		case conn, ok := <-connectionChannel:
			if !ok {
				return
			}
			defer conn.Close()
		}
	}
}

type DefaultWorkerPool struct {
	connectionChannel chan net.Conn
}

func NewDefaultWorkerPool() WorkerPool {
	return &DefaultWorkerPool{
		connectionChannel: nil,
	}
}

func (p *DefaultWorkerPool) StartWorkers(ctx context.Context, workers ...uint) {
	if p.connectionChannel != nil {
		close(p.connectionChannel)
	}

	var workersCount uint = defaultWorkersCount
	if len(workers) > 0 {
		workersCount = workers[0]
	}

	p.connectionChannel = make(chan net.Conn, workersCount)

	for i := uint(0); i < workersCount; i++ {
		go newWorker(ctx, p.connectionChannel)
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
