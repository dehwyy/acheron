package packet

import (
	"encoding/binary"
	"io"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/xdptypes"
)

func CreatePacketFromReader(r io.Reader) (*Packet, error) {
	p, err := newPacketWithMetadataFromReader(r)
	if err != nil {
		return nil, err
	}

	// Parse headers
	headersBuffer := make([]byte, p.HeadersLen)
	if _, err = io.ReadFull(r, headersBuffer); err != nil {
		return nil, err
	}
	p.Headers, err = headersFromBytes(headersBuffer)
	if err != nil {
		return nil, err
	}

	// Parse payload
	payloadBuffer := make([]byte, p.PayloadLen)
	if _, err = io.ReadFull(r, payloadBuffer); err != nil {
		return nil, err
	}
	p.Payload, err = payloadFromBytes(payloadBuffer)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func newPacketWithMetadataFromReader(r io.Reader) (*Packet, error) {
	var p Packet

	metadata := make([]byte, 1+1+2+4)
	_, err := io.ReadFull(r, metadata)
	if err != nil {
		log.Logger.Error().Msgf("Failed to read metadata: %v", err)
		return nil, err
	}

	p.Version = metadata[0]
	p.PacketType = xdptypes.PacketType(metadata[1])
	p.HeadersLen = binary.BigEndian.Uint16(metadata[2:4])
	p.PayloadLen = binary.BigEndian.Uint32(metadata[4:8])

	return &p, nil
}
