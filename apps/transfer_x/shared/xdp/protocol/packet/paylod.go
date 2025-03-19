package packet

import (
	"encoding/binary"
	"reflect"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/log"
	xd "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet/xdptypes"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

const (
	offsetPayloadKeyLen   uint32 = 0
	offsetPayloadValueLen uint32 = 1
	offsetPayloadDataType uint32 = 5
	offsetPayloadData     uint32 = 6
)

func payloadFromBytes(b []byte) ([]Field, error) {
	payload := make([]Field, 0, 4)
	size := uint32(len(b))
	var offset uint32

	for offset < size {
		keyLen := uint32(b[offset+offsetPayloadKeyLen])
		valueLen := binary.BigEndian.Uint32(b[offset+offsetPayloadValueLen : offset+offsetPayloadDataType])

		newOffset := offset + offsetPayloadData + keyLen + valueLen
		if newOffset > size {
			log.Logger.Warn().Msgf("Limit exceeded (Payload): %d/%d", newOffset, size)
			break
		}

		payload = append(payload, Field{
			KeyLen:   b[offset],
			DataType: xd.PayloadDataType(b[offset+offsetPayloadDataType]),
			ValueLen: valueLen,
			Key:      b[offset+offsetPayloadData : offset+offsetPayloadData+keyLen],
			Value:    b[offset+offsetPayloadData+keyLen : offset+offsetPayloadData+keyLen+valueLen],
		})

		offset = newOffset
	}

	return payload, nil
}

func createParsingCallback(field Field) func(callback func([]byte) any, typeSize uint8) any {
	return func(callback func([]byte) any, typeSize uint8) any {
		if !xd.IsArray(field.DataType) {
			return callback(field.Value)
		}

		values := make([]any, len(field.Value)/int(typeSize))

		for i := 0; i < len(field.Value)-1; i += int(typeSize) {
			values[i/int(typeSize)] = callback(field.Value[i : i+int(typeSize)])
		}

		return values
	}
}

func newPayloadFromReflect(payload []Field, reflectType reflect.Type) reflect.Value {
	reflectedPayload := reflect.New(reflectType)

	for _, field := range payload {
		var value any
		key := string(field.Key)

		fromCallback := createParsingCallback(field)

		switch field.DataType {
		case xd.U8:
			value = fromCallback(func(b []byte) any { return b[0] }, 1)
		case xd.U16:
			value = fromCallback(func(b []byte) any { return binary.BigEndian.Uint16(b) }, 2)
		case xd.U32:
			value = fromCallback(func(b []byte) any { return binary.BigEndian.Uint32(b) }, 4)
		case xd.U64:
			value = fromCallback(func(b []byte) any { return binary.BigEndian.Uint64(b) }, 8)
		case xd.I8:
			value = fromCallback(func(b []byte) any { return int(b[0]) }, 1)
		case xd.I16:
			value = fromCallback(func(b []byte) any { return int16(binary.BigEndian.Uint16(b)) }, 2)
		case xd.I32:
			value = fromCallback(func(b []byte) any { return int32(binary.BigEndian.Uint32(b)) }, 4)
		case xd.I64:
			value = fromCallback(func(b []byte) any { return int64(binary.BigEndian.Uint64(b)) }, 8)
		case xd.F32:
			value = fromCallback(func(b []byte) any { return float32(binary.BigEndian.Uint32(b)) }, 4)
		case xd.F64:
			value = fromCallback(func(b []byte) any { return float64(binary.BigEndian.Uint64(b)) }, 8)
		case xd.Boolean:
			value = fromCallback(func(b []byte) any { return b[0] != 0 }, 1)
		case xd.String:
			value = string(field.Value)
		case xd.StringArray:
			// TODO
		case xd.Nested:
			data, err := payloadFromBytes(field.Value)
			if err != nil {

			}

			value = newPayloadFromReflect(data, reflectedPayload.FieldByName(key).Type())

		case xd.ArrayMask:
			log.Logger.Error().Msgf("<Mask> cannot be data type")
		default:
			log.Logger.Error().Msgf("Unknown type: %v", field.DataType)
			return reflectedPayload
		}

		reflectedPayload.FieldByName(key).Set(reflect.ValueOf(value))
	}

	return reflectedPayload
}

func CreatePayload[T types.Payload](rawPayload []Field) T {
	var payload T
	reflectedPayload := reflect.ValueOf(payload)

	payload, ok := newPayloadFromReflect(rawPayload, reflectedPayload.Type()).Interface().(T)
	if !ok {
		log.Logger.Error().Msgf("Failed to create payload: %v", payload)
	}

	return payload
}
