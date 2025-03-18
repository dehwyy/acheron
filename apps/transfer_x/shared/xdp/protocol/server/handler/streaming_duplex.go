package handler

import t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"

type StreamingDuplexHandler[Req t.StreamPayload] struct {
	Handler func(rx <-chan t.StreamRequest[Req], tx chan<- t.StreamResponse[t.StreamPayload]) error
}

func NewStreamingDuplexHandler[Req t.StreamPayload](h func(<-chan t.StreamRequest[Req], chan<- t.StreamResponse[t.StreamPayload]) error) StreamingHandler[t.StreamPayload] {
	return &StreamingDuplexHandler[Req]{Handler: h}
}

func (h *StreamingDuplexHandler[Req]) Handle(rx <-chan t.StreamRequest[t.StreamPayload], tx chan<- t.StreamResponse[t.StreamPayload]) error {
	rxPiped := make(chan t.StreamRequest[Req], cap(rx))

	go func() {
		for req := range rx {
			rxPiped <- req.(t.StreamRequest[Req]) // TODO
		}
	}()

	return h.Handler(rxPiped, tx)
}
