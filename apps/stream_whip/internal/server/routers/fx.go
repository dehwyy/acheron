package routers

import (
	"github.com/dehwyy/acheron/apps/stream_whip/internal/rtc"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"go.uber.org/fx"
)

type WhipWhepRouterParams struct {
	fx.In

	Log logg.Logger
}

func NewWhipWhepRouterFx(params WhipWhepRouterParams) *WhipWhepRouter {
	api, err := rtc.NewAPI()
	if err != nil {
		panic(err)
	}

	return &WhipWhepRouter{
		log: params.Log,
		api: api,
	}
}
