package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
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

type RawPayload struct {
	fields []rawPayloadField
}

type rawPayloadField struct {
	Key      []byte
	Value    []byte
	KeyLen   byte
	DataType xd.PayloadDataType
	ValueLen uint32
}

func NewRawPayload(b []byte) (*RawPayload, error) {
	payload := make([]rawPayloadField, 0, 4)
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

		payload = append(payload, rawPayloadField{
			KeyLen:   b[offset],
			DataType: xd.PayloadDataType(b[offset+offsetPayloadDataType]),
			ValueLen: valueLen,
			Key:      b[offset+offsetPayloadData : offset+offsetPayloadData+keyLen],
			Value:    b[offset+offsetPayloadData+keyLen : offset+offsetPayloadData+keyLen+valueLen],
		})

		offset = newOffset
	}

	return &RawPayload{payload}, nil
}

func (*RawPayload) createParsingCallback(v []byte, t xd.PayloadDataType) func(callback func([]byte) any, size uint8) any {
	return func(callback func([]byte) any, sz uint8) any {
		size := int(sz)
		if !xd.IsArray(t) {
			return callback(v)
		}

		values := make([]any, len(v)/size)

		// ? Maybe err
		for i := 0; i < len(v)-1; i += size {
			values[i/size] = callback(v[i : i+size])
		}

		return values
	}
}

func (raw *RawPayload) ToPayloadReflected(reflectType reflect.Type) reflect.Value {
	payload := reflect.New(reflectType).Elem()

	for _, f := range raw.fields {
		var value any
		key := string(f.Key)

		fromCallback := raw.createParsingCallback(f.Value, f.DataType)

		switch f.DataType {
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
			value = string(f.Value)
		case xd.StringArray:
			// TODO
		case xd.Nested:
			nestedRaw, err := NewRawPayload(f.Value)
			if err != nil {
				log.Logger.Error().Msgf("Failed to create nested payload: %v", err)
				continue
			}
			value = nestedRaw.ToPayloadReflected(payload.FieldByName(key).Type())
		case xd.ArrayMask:
			log.Logger.Error().Msgf("<Mask> cannot be data type")
		default:
			log.Logger.Error().Msgf("Unknown type: %v", f.DataType)
			return payload
		}

		payload.FieldByName(key).Set(reflect.ValueOf(value))
	}

	return payload
}

func PayloadFromRaw[T types.Payload](rawPayload *RawPayload) (*T, error) {
	var payload T

	payload, ok := rawPayload.ToPayloadReflected(reflect.TypeOf(payload)).Interface().(T)
	if !ok {
		log.Logger.Error().Msgf("Failed to create payload: %v", payload)
		return nil, errors.New("Failed to create payload")
	}

	return &payload, nil
}

func PayloadFromBytes[T types.Payload](b []byte) (*T, error) {
	var payload T

	rawPayload, err := NewRawPayload(b)
	if err != nil {
		return nil, err
	}

	payload, ok := rawPayload.ToPayloadReflected(reflect.TypeOf(payload)).Interface().(T)
	if !ok {
		log.Logger.Error().Msgf("Failed to create payload: %v", payload)
		return nil, errors.New("Failed to create payload")
	}

	return &payload, nil
}

func PayloadToBytes[T types.Payload](payload *T) ([]byte, error) {
	reflectPayload := reflect.ValueOf(payload).Elem()
	reflectPayloadType := reflect.TypeOf(payload).Elem()

	var buf bytes.Buffer
	for i := range reflectPayload.NumField() {
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
