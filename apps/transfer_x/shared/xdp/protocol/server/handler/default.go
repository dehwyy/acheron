package handler

import t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"

type DefaultHandler[T t.Payload] struct {
	Handler func(t.Request[T]) error
}

func NewDefaultHandler[T t.Payload](handler func(t.Request[T]) error) Handler[t.Payload] {
	return &DefaultHandler[T]{Handler: handler}
}

func (h *DefaultHandler[T]) Handle(req t.Request[t.Payload]) error {
	return h.Handler(req.(t.Request[T]))
}
