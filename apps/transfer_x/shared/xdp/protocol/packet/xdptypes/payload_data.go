package xdptypes

import (
	"reflect"
)

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

func FromReflectKind(k reflect.Kind) PayloadDataType {
	switch k {
	case reflect.Uint8:
		return U8
	case reflect.Uint16:
		return U16
	case reflect.Uint32:
		return U32
	case reflect.Uint64:
		return U64
	case reflect.Int8:
		return I8
	case reflect.Int16:
		return I16
	case reflect.Int32:
		return I32
	case reflect.Int64:
		return I64
	case reflect.Float32:
		return F32
	case reflect.Float64:
		return F64
	case reflect.Bool:
		return Boolean
	case reflect.String:
		return String
	default:
		return 0
	}
}

func IsArray(t PayloadDataType) bool {
	return t&ArrayMask>>7 == 1
}
