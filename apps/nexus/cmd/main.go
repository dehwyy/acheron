package main

import (
	"github.com/dehwyy/acheron/apps/nexus/internal/gql"
	"github.com/dehwyy/acheron/apps/nexus/internal/gql/gqlgen"
	"github.com/dehwyy/acheron/apps/nexus/internal/gql/resolvers"
	"github.com/dehwyy/acheron/apps/nexus/internal/server"
	"github.com/dehwyy/acheron/libraries/go/config"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New(config.ConfigConstructorParams{}),
			logg.New(logg.Opts{
				ServiceName: "nexus",
			}),
			fx.Annotate(resolvers.New, fx.As(new(gqlgen.ResolverRoot))),
			server.NewFx,
		),
		fx.Invoke(gql.NewFx),
	).Run()
}
