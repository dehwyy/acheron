package connection

import "net"

type ConnectionHandler struct{}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{}
}

func (c *ConnectionHandler) HandleConnection(conn net.Conn) error {
	return nil
}
