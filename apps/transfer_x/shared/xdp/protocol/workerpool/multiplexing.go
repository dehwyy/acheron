package workerpool

import "net"

type key uint
type MultiplexingWorkerPool struct {
	connectionChannels map[key]chan net.Conn // Is it So? Should *Packet be passed instead?
}
