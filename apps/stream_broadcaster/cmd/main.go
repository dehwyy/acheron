package main

import (
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/gqlgen"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/resolvers"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/server"
	"github.com/dehwyy/mugen/libraries/go/config"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New(config.Opts{}),
			logg.New(logg.Opts{
				ServiceName: "stream_broadcaster",
			}),
			fx.Annotate(resolvers.New, fx.As(new(gqlgen.ResolverRoot))),
			server.NewFx,
		),
		fx.Invoke(gql.NewFx),
	).Run()
}
