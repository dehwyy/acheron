package connection

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/xdptypes"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
)

type ConnectionHandler struct {
	router router.Router
}

func NewConnectionHandler(r router.Router) *ConnectionHandler {
	return &ConnectionHandler{router: r}
}

func (c *ConnectionHandler) HandleConnection(conn net.Conn) error {
	defer conn.Close()

	var p packet.Packet
	// Reading metadata
	metadata := make([]byte, 1+1+2+4)
	_, err := io.ReadFull(conn, metadata)
	if err != nil {
		log.Logger.Error().Msgf("Failed to read metadata: %v", err)
		return err
	}

	_ = fillPacketMetadata(metadata, &p)

	headers := make([]byte, p.HeadersLen)

	_, err = io.ReadFull(conn, headers)
	if err != nil {
		log.Logger.Error().Msgf("Failed to read headers: %v", err)
		return err
	}

	return nil
}

func fillPacketMetadata(metadata []byte, p *packet.Packet) error {
	p.Version = metadata[0]
	p.PacketType = xdptypes.PacketType(metadata[1])
	p.HeadersLen = binary.BigEndian.Uint16(metadata[2:4])
	p.PayloadLen = binary.BigEndian.Uint32(metadata[4:8])

	return nil
}
