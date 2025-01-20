package routers

import (
	"github.com/dehwyy/mugen/libraries/go/logg"
	"go.uber.org/fx"
)

type WhipWhepRouterParams struct {
	fx.In

	Log logg.Logger
}

func NewWhipWhepRouterFx(params WhipWhepRouterParams) *WhipWhepRouter {
	return &WhipWhepRouter{
		log: params.Log,
	}
}
