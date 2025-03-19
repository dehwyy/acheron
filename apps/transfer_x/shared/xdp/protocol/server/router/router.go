package router

import (
	h "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/handler"
	t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

type Router interface {
	ReadableRouter
	AddRoute(route string, handler h.Handler[t.Payload])
	AddStreamingRoute(route string, handler h.StreamingHandler[t.StreamPayload])
	Mount(baseRoute string, router Router)
}

type RouteType uint8

const (
	ClassicRoute   RouteType = 1
	StreamingRoute RouteType = 2
)

type ReadableRouter interface {
	GetRouteType(route string) RouteType
	GetClassicRoute(route string) h.Handler[t.Payload]
	GetStreamingRoute(route string) h.StreamingHandler[t.Payload]
}
