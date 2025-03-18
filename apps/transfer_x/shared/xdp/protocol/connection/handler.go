package connection

import (
	"net"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
)

type ConnectionHandler struct {
	router router.Router
}

func NewConnectionHandler(r router.Router) *ConnectionHandler {
	return &ConnectionHandler{router: r}
}

func (c *ConnectionHandler) HandleConnection(conn net.Conn) error {
	return nil
}
