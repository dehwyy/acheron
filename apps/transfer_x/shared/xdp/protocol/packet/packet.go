package packet

import "github.com/dehwyy/acheron/apps/stream_x/shared/xdp/protocol/packet/xdptypes"

type Packet struct {
	version      byte
	length       uint32
	headerLength uint16
	headers      []any // TODO
	payload      []Payload
}

type Payload struct {
	keyLength uint16
	key       []byte
	dataType  xdptypes.PayloadDataType
	data      []byte
}
