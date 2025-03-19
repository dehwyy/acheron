package packet

import (
	"encoding/binary"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
)

// ! - Necessary headers
// ? - Optional headers (depends of `packetType“)
const (
	HeaderRoute    = "route"     // !
	HeaderPacketID = "packet-id" // !
	HeaderStreamID = "stream-id" // ?
)

const (
	offsetHeaderKeyLen   uint16 = 0
	offsetHeaderValueLen uint16 = 1
	offsetHeaderData     uint16 = 3
)

func headersFromBytes(b []byte) ([]Header, error) {
	var headers []Header
	size := uint16(len(b))
	var offset uint16

	for offset < size {
		keyLen := uint16(b[offset+offsetHeaderKeyLen])
		valueLen := binary.BigEndian.Uint16(b[offset+offsetHeaderValueLen : offset+offsetHeaderData])

		newOffset := offset + offsetHeaderData + keyLen + valueLen
		if newOffset > size {
			log.Logger.Warn().Msgf("Limit exceeded (Headers): %d/%d", newOffset, size)
			return headers, nil
		}

		headers = append(headers, Header{
			KeyLen:   b[offset],
			ValueLen: valueLen,
			Key:      b[offset+offsetHeaderData : offset+offsetHeaderData+keyLen],
			Value:    b[offset+offsetHeaderData+keyLen : offset+offsetHeaderData+keyLen+valueLen],
		})

		offset = newOffset
	}

	return headers, nil
}

func CreateHeadersMap(headers []Header) map[string]string {
	h := make(map[string]string)
	for _, header := range headers {
		h[string(header.Key)] = string(header.Value)
	}

	return h
}
