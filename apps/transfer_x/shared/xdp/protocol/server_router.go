package xdp

import (
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/handler"
	"github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/router"
)

func NewRouter() router.Router {
	return &router.DefaultRouter{}
}

// Wrapper function allows to add static typed payload to endpoints.
func AddRoute[T Payload](r router.Router, route string, h func(Request[T]) error) {
	r.AddRoute(route, handler.NewDefaultHandler(h))
}

func AddStreamingRoute[Req StreamPayload](r router.Router, route string, h func(rx <-chan StreamRequest[Req], tx chan<- StreamResponse[StreamPayload]) error) {
	r.AddStreamingRoute(route, handler.NewStreamingDuplexHandler(h))
}
