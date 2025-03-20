package packet_test

import (
	"fmt"
	"testing"

	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet"
)

type P struct {
	Unsigned8  uint8
	Unsigned16 uint16
	Unsigned32 uint32
	Unsigned64 uint64
	Signed8    int8
	Signed16   int16
	Signed32   int32
	Signed64   int64
	Float32    float32
	Float64    float64
	Boolean    bool
	SomeString string
}

func TestEncodeDecodeBinary(t *testing.T) {
	p := P{
		Unsigned8:  1,
		Unsigned16: 2,
		Unsigned32: 3,
		Unsigned64: 4,
		Signed8:    -5,
		Signed16:   -6,
		Signed32:   -7,
		Signed64:   -8,
		Float32:    9.0,
		Float64:    10.0,
		Boolean:    true,
		SomeString: "hello typeshit ts pmo",
	}
	b, err := packet.PayloadToBytes(p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := packet.RawPayloadFromBytes(b)
	if err != nil {
		t.Fatal(err)
	}

	parsed := packet.PayloadFromRaw[P](r)

	fmt.Printf("initial: %+v\n", p)
	fmt.Printf("parsed : %+v\n", parsed)
}
