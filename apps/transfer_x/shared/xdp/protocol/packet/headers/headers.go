package headers

import (
	"encoding/binary"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
)

type RawHeaders struct {
	headers []rawHeader
}

type rawHeader struct {
	Value    []byte
	Key      []byte
	ValueLen uint16
	KeyLen   byte
}

const (
	offsetHeaderKeyLen   uint16 = 0
	offsetHeaderValueLen uint16 = 1
	offsetHeaderData     uint16 = 3
)

func NewRawHeaders(b []byte) (*RawHeaders, error) {
	var headers []rawHeader
	size := uint16(len(b))
	var offset uint16

	for offset < size {
		keyLen := uint16(b[offset+offsetHeaderKeyLen])
		valueLen := binary.BigEndian.Uint16(b[offset+offsetHeaderValueLen : offset+offsetHeaderData])

		newOffset := offset + offsetHeaderData + keyLen + valueLen
		if newOffset > size {
			log.Logger.Warn().Msgf("Limit exceeded (Headers): %d/%d", newOffset, size)
			break
		}

		headers = append(headers, rawHeader{
			KeyLen:   b[offset],
			ValueLen: valueLen,
			Key:      b[offset+offsetHeaderData : offset+offsetHeaderData+keyLen],
			Value:    b[offset+offsetHeaderData+keyLen : offset+offsetHeaderData+keyLen+valueLen],
		})

		offset = newOffset
	}

	return &RawHeaders{headers}, nil
}

func (raw *RawHeaders) ToMap() map[string]string {
	h := make(map[string]string)
	for _, header := range raw.headers {
		h[string(header.Key)] = string(header.Value)
	}

	return h
}
