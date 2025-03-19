package xdptypes

type PayloadDataType byte

const (
	U8  PayloadDataType = 1
	U16 PayloadDataType = 2
	U32 PayloadDataType = 3
	U64 PayloadDataType = 4

	I8  PayloadDataType = 5
	I16 PayloadDataType = 6
	I32 PayloadDataType = 7
	I64 PayloadDataType = 8

	F32 PayloadDataType = 9
	F64 PayloadDataType = 10

	Boolean PayloadDataType = 11

	String      PayloadDataType = 1 << 5      // 32
	StringArray PayloadDataType = 1<<6 | 1<<5 // hz

	ArrayMask PayloadDataType = 1 << 7   // 128
	Nested    PayloadDataType = 1<<8 - 1 // 255
)

func IsArray(t PayloadDataType) bool {
	return t&ArrayMask>>7 == 1
}
