package main

import (
	"github.com/dehwyy/mugen/apps/stream_whip/internal/server"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/server/routers"
	"github.com/dehwyy/mugen/libraries/go/config"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New(config.ConfigConstructorParams{}),
			logg.New(logg.Opts{
				ServiceName: "stream_whip-whep",
			}),
		),
		fx.Provide(routers.NewWhipWhepRouterFx),
		fx.Invoke(server.NewFx),
	).Run()
}
