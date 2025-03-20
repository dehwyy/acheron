package connection

import (
	"net"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/headers"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

type ConnectionHandler struct {
	router router.ReadableRouter
}

func NewConnectionHandler(r router.ReadableRouter) *ConnectionHandler {
	return &ConnectionHandler{router: r}
}

func (c *ConnectionHandler) HandleConnection(conn net.Conn) error {
	defer conn.Close()

	p, err := packet.NewPacket(conn)
	if err != nil {
		return err
	}

	h := p.Headers.ToMap()
	route, ok := h[headers.HeaderRoute]
	if !ok {
		log.Logger.Error().Msgf("No route header in packet: %v", p.Headers)
		return nil
	}

	routeType := c.router.GetRouteType(route)
	switch routeType {
	case router.ClassicRoute:
		err = c.router.GetClassicRoute(route).Handle(types.NewRequest(p.Payload))
		if err != nil {
			log.Logger.Error().Msgf("Failed to handle classic route: %v", err)
			return err
		}
	case router.StreamingRoute:
		// TODO
		// c.router.GetStreamingRoute(route).Handle(p.Payload)
	}

	return nil
}
