package workerpool

import (
	"context"
	"net"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
)

type WorkerPool interface {
	StartWorkers(ctx context.Context, r router.ReadableRouter, workers ...uint)
	Stop()

	QueueConnection(net.Conn) <-chan error
}

func NewWorkerPool() *DefaultWorkerPool {
	return &DefaultWorkerPool{connectionChannel: nil}
}
