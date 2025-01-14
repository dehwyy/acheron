package main

import (
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/server"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logg.New(logg.Opts{
				ServiceName: "stream_broadcaster",
			}),
			server.NewFx,
		),
		fx.Invoke(gql.NewFx),
	).Run()
}
