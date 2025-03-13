package workerpool

import (
	"context"
	"net"
)

type WorkerPool interface {
	StartWorkers(ctx context.Context, workers ...uint)
	Stop()

	QueueConnection(net.Conn) <-chan error
}
