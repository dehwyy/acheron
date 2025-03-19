package xdptypes

type PacketType byte

const (
	PacketTypeRequest        PacketType = 0x01
	PacketTypeResponse       PacketType = 0x02
	PacketTypeStreamRequest  PacketType = 0x03
	PacketTypeStreamResponse PacketType = 0x04
)
