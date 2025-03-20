package packet

import (
	"encoding/binary"
	"io"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/headers"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/xdptypes"
)

const (
	offsetMetaVersion    uint8 = 0
	offsetMetaPacketType uint8 = 1
	offsetMetaHeadersLen uint8 = 2
	offsetMetaPayloadLen uint8 = 4
	offsetMetaData       uint8 = 8
)

func NewPacket(r io.Reader) (*Packet, error) {
	p := new(Packet)

	metadata := make([]byte, offsetMetaData)
	_, err := io.ReadFull(r, metadata)
	if err != nil {
		log.Logger.Error().Msgf("Failed to read metadata: %v", err)
		return nil, err
	}

	p.ProtocolVersion = metadata[offsetMetaVersion]
	p.PacketType = xdptypes.PacketType(metadata[offsetMetaPacketType])
	headersLen := binary.BigEndian.Uint16(metadata[offsetMetaHeadersLen:offsetMetaPayloadLen])
	payloadLen := binary.BigEndian.Uint32(metadata[offsetMetaPayloadLen:offsetMetaData])

	// Parse headers
	headersBuffer := make([]byte, headersLen)
	if _, err = io.ReadFull(r, headersBuffer); err != nil {
		return nil, err
	}

	p.Headers, err = headers.NewRawHeaders(headersBuffer)
	if err != nil {
		return nil, err
	}

	// Parse payload
	payloadBuffer := make([]byte, payloadLen)
	if _, err = io.ReadFull(r, payloadBuffer); err != nil {
		return nil, err
	}
	p.Payload, err = NewRawPayload(payloadBuffer)
	if err != nil {
		return nil, err
	}

	return p, nil
}
