package handler

import (
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet"
	t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

type DefaultHandler[T t.Payload] struct {
	Handler func(t.Request[T]) error
}

func NewDefaultHandler[T t.Payload](handler func(t.Request[T]) error) Handler[t.Payload] {
	return &DefaultHandler[T]{Handler: handler}
}

func (h *DefaultHandler[T]) Handle(req t.Request[*packet.RawPayload]) error {
	p, err := packet.PayloadFromRaw[T](req.Get())
	if err != nil {
		return err
	}

	req2 := t.NewRequest(*p)

	return h.Handler(req2)
}
