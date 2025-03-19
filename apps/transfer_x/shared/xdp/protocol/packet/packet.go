package packet

import "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/xdptypes"

type Packet struct {
	Version    byte
	PacketType xdptypes.PacketType
	HeadersLen uint16
	PayloadLen uint32
	Headers    []Header
	Payload    []Field
}

type Header struct {
	KeyLen   byte
	ValueLen uint16
	Key      []byte
	Value    []byte
}

type Field struct {
	KeyLen   byte
	DataType xdptypes.PayloadDataType
	ValueLen uint32
	Key      []byte
	Value    []byte
}
