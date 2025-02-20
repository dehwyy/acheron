package main

import (
	"github.com/dehwyy/acheron/apps/stream_broadcaster/internal/repos"
	"github.com/dehwyy/acheron/apps/stream_broadcaster/internal/server"
	"github.com/dehwyy/acheron/apps/stream_broadcaster/internal/server/routers"
	"github.com/dehwyy/acheron/libraries/go/config"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New(config.ConfigConstructorParams{}),
			logg.New(logg.Opts{
				ServiceName: "stream_broadcaster",
			}),
		),
		// repositories
		fx.Provide(repos.NewFileRepoFx),
		// routers
		fx.Provide(routers.NewPlaylistRouterFx),
		fx.Invoke(server.NewFx),
	).Run()
}
