package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
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

func RawPayloadFromBytes(b []byte) (RawPayload, error) {
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

func createParsingCallback(field Field) func(callback func([]byte) any, size uint8) any {
	return func(callback func([]byte) any, sz uint8) any {
		size := int(sz)
		if !xd.IsArray(field.DataType) {
			return callback(field.Value)
		}

		values := make([]any, len(field.Value)/size)

		for i := 0; i < len(field.Value)-1; i += size {
			values[i/size] = callback(field.Value[i : i+size])
		}

		return values
	}
}

func payloadFromRawReflected(rawPayload RawPayload, reflectType reflect.Type) reflect.Value {
	payload := reflect.New(reflectType)
	if payload.Kind() == reflect.Ptr {
		payload = payload.Elem()
	}

	for _, field := range rawPayload {
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
			value = fromCallback(func(b []byte) any { return int8(b[0]) }, 1)
		case xd.I16:
			value = fromCallback(func(b []byte) any { return int16(binary.BigEndian.Uint16(b)) }, 2)
		case xd.I32:
			value = fromCallback(func(b []byte) any { return int32(binary.BigEndian.Uint32(b)) }, 4)
		case xd.I64:
			value = fromCallback(func(b []byte) any { return int64(binary.BigEndian.Uint64(b)) }, 8)
		case xd.F32:
			value = fromCallback(func(b []byte) any { return math.Float32frombits(binary.BigEndian.Uint32(b)) }, 4)
		case xd.F64:
			value = fromCallback(func(b []byte) any { return math.Float64frombits(binary.BigEndian.Uint64(b)) }, 8)
		case xd.Boolean:
			value = fromCallback(func(b []byte) any { return b[0] != 0 }, 1)
		case xd.String:
			value = string(field.Value)
		case xd.StringArray:
			// TODO
		case xd.Nested:
			data, err := RawPayloadFromBytes(field.Value)
			if err != nil {

			}

			value = payloadFromRawReflected(data, payload.FieldByName(key).Type())

		case xd.ArrayMask:
			log.Logger.Error().Msgf("<Mask> cannot be data type")
		default:
			log.Logger.Error().Msgf("Unknown type: %v", field.DataType)
			return payload
		}

		payload.FieldByName(key).Set(reflect.ValueOf(value))
	}

	return payload
}

func PayloadFromRaw[T types.Payload](rawPayload []Field) T {
	var payload T

	payload, ok := payloadFromRawReflected(rawPayload, reflect.TypeOf(payload)).Interface().(T)
	if !ok {
		log.Logger.Error().Msgf("Failed to create payload: %v", payload)
	}

	return payload
}

func PayloadToBytes[T types.Payload](payload T) ([]byte, error) {
	reflectPayload := reflect.ValueOf(payload)
	reflectPayloadType := reflect.TypeOf(payload)

	var buf bytes.Buffer
	for i := 0; i < reflectPayload.NumField(); i++ {
		field := reflectPayload.Field(i)
		fieldType := reflectPayloadType.Field(i)

		key := []byte(fieldType.Name)
		value := field.Interface()

		var size uint32

		switch field.Kind() {
		case reflect.String:
			size = uint32(len([]byte((field.String()))))
		default:
			size = uint32(field.Type().Size())
		}

		fmt.Println("size: ", size, " with kind: ", field.Kind())

		if err := binary.Write(&buf, binary.BigEndian, byte(len(key))); err != nil {
			return nil, err
		}
		if err := binary.Write(&buf, binary.BigEndian, size); err != nil {
			return nil, err
		}
		if err := binary.Write(&buf, binary.BigEndian, xd.FromReflectKind(field.Kind())); err != nil {
			return nil, err
		}

		if err := binary.Write(&buf, binary.BigEndian, key); err != nil {
			return nil, err
		}

		switch field.Kind() {
		case reflect.String:
			if _, err := buf.WriteString(field.String()); err != nil {
				return nil, err
			}
		default:
			if err := binary.Write(&buf, binary.BigEndian, value); err != nil {
				return nil, err
			}
		}

	}

	return buf.Bytes(), nil
}

func _payloadToRawReflected(payload reflect.Value) RawPayload {
	rawPayload := make(RawPayload, 0)

	for i := 0; i < payload.NumField(); i++ {
		field := payload.Field(i)

		key := []byte(field.Type().Name())
		var value []byte

		switch field.Kind() {
		case reflect.Uint8:
			value = make([]byte, 1)
			value[0] = uint8(field.Uint())
		case reflect.Uint16:
			value = make([]byte, 2)
			binary.BigEndian.PutUint16(value, uint16(field.Uint()))
		case reflect.Uint32:
			value = make([]byte, 4)
			binary.BigEndian.PutUint32(value, uint32(field.Uint()))
		case reflect.Uint64:
			value = make([]byte, 8)
			binary.BigEndian.PutUint64(value, uint64(field.Uint()))
		case reflect.Int8:
			value = make([]byte, 1)
			value[0] = uint8(field.Int())
		case reflect.Int16:
			value = make([]byte, 2)
			binary.BigEndian.PutUint16(value, uint16(field.Int()))
		case reflect.Int32:
			value = make([]byte, 4)
			binary.BigEndian.PutUint32(value, uint32(field.Int()))
		case reflect.Int64:
			value = make([]byte, 8)
			binary.BigEndian.PutUint64(value, uint64(field.Int()))
		case reflect.Float32:
			value = make([]byte, 4)
			binary.BigEndian.PutUint32(value, math.Float32bits(float32(field.Float())))
		case reflect.Float64:
			value = make([]byte, 8)
			binary.BigEndian.PutUint64(value, math.Float64bits(field.Float()))
		case reflect.Bool:
			value = make([]byte, 1)
			if field.Bool() {
				value[0] = 1
			}
		case reflect.String:
			value = []byte(field.String())
		case reflect.Slice:
			// TODO
		}

		rawPayload = append(rawPayload, Field{
			KeyLen:   byte(len(key)),
			DataType: xd.FromReflectKind(field.Kind()),
			ValueLen: uint32(len(value)),
			Key:      key,
			Value:    value,
		})
	}

	return nil
}

func _PayloadToRaw[T types.Payload](payload T) (RawPayload, error) {

	rawPayload := _payloadToRawReflected(reflect.ValueOf(payload))

	return rawPayload, nil
}
