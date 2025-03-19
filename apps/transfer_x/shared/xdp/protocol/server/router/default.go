package router

import (
	h "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/server/handler"
	t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"
)

type DefaultRouter struct {
	routes       map[string]h.Handler[t.Payload]
	streamRoutes map[string]h.StreamingHandler[t.StreamPayload]
}

// `Router` interface

func (r *DefaultRouter) AddRoute(route string, handler h.Handler[t.Payload]) {
	r.routes[route] = handler
}

func (r *DefaultRouter) AddStreamingRoute(route string, handler h.StreamingHandler[t.StreamPayload]) {
	r.streamRoutes[route] = handler
}

func (r *DefaultRouter) Mount(baseRoute string, router Router) {} // TODO

// `ReadableRouter` interface

func (r *DefaultRouter) GetRouteType(route string) RouteType {
	return ClassicRoute // TODO
}

func (r *DefaultRouter) GetClassicRoute(route string) h.Handler[t.Payload] {
	return r.routes[route]
}

func (r *DefaultRouter) GetStreamingRoute(route string) h.StreamingHandler[t.Payload] {
	return nil // TODO
}
