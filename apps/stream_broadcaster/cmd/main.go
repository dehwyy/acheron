package main

import (
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(server.New),
		fx.Invoke(gql.New),
	).Run()
}
