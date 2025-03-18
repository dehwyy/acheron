package router

import (
	h "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/handler"
	t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

type Router interface {
	AddRoute(route string, handler h.Handler[t.Payload])
	AddStreamingRoute(route string, handler h.StreamingHandler[t.StreamPayload])
	Mount(baseRoute string, router Router)
}

type ReadableRouter interface {
	GetRoute(route string) h.Handler[t.Payload]
}
