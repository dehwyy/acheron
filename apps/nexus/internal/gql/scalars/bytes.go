package scalars

import (
	"fmt"
	"io"
)

type Bytes []byte

func (b *Bytes) UnmarshalGQL(v any) error {
	bytes, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("cannot cast %v to []byte", v)
	}

	copy(*b, bytes)
	return nil
}

func (b Bytes) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%s"`, b)
}
