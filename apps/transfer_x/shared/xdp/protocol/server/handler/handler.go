package handler

import (
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/packet"
	t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

type Handler[P t.Payload] interface {
	Handle(t.Request[[]packet.Field]) error // TODO: custom error
}

type StreamingHandler[Req t.StreamPayload] interface {
	Handle(rx <-chan t.StreamRequest[Req], tx chan<- t.StreamResponse[t.StreamPayload]) error
}
