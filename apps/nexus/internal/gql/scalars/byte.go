package scalars

import (
	"fmt"
	"io"
)

type Byte byte

func (b *Byte) UnmarshalGQL(v any) error {
	castedValue, ok := v.(byte)
	if !ok {
		return fmt.Errorf("cannot cast %v to []byte", v)
	}

	*b = Byte(castedValue)
	return nil
}

func (b Byte) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%c"`, b)
}
