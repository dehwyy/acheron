package packet

import (
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/headers"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/xdptypes"
)

type Packet struct {
	Headers         *headers.RawHeaders
	Payload         *RawPayload
	ProtocolVersion byte
	PacketType      xdptypes.PacketType
}
