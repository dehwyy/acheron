package scalars

import (
	"fmt"
	"io"
)

type Bytes []byte

func (b *Bytes) UnmarshalGQL(v interface{}) error {
	bytes, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("cannot cast %v to []byte", v)
	}

	*b = bytes
	return nil
}

func (b Bytes) MarshalGQL(w io.Writer) {
	w.Write(b)
}
