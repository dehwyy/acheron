package server

import (
	"context"

	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/server/routers"
	"github.com/dehwyy/mugen/libraries/go/config"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	LC     fx.Lifecycle
	Log    logg.Logger
	Config config.Config

	PlaylistRouter *routers.PlaylistRouter
}

func NewFx(opts Opts) *Server {
	r := &Server{
		gin.New(),
	}

	r.Use(
		cors.New(
			cors.Config{
				AllowAllOrigins:  true,
				AllowMethods:     []string{"*"},
				AllowHeaders:     []string{"*"},
				ExposeHeaders:    []string{"*"},
				AllowCredentials: true,
			},
		),
	)

	v1 := r.Group("/api/v1")

	opts.PlaylistRouter.RegisterRoutes(v1)

	opts.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				opts.Log.Debug().Msgf("Provided config: %s", opts.Config)
				opts.Log.Info().Msg("Starting server...")

				r.Start(ctx, opts.Config.Addr().Ports.StreamBroadcasterPort)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			opts.Log.Info().Msg("Stopping server...")
			return r.Stop(ctx)
		},
	})

	return r
}
