package server

import (
	"context"

	"github.com/dehwyy/mugen/libraries/go/config"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	LC     fx.Lifecycle
	Log    logg.Logger
	Config config.Config
}

func NewFx(opts Opts) *Server {
	r := &Server{
		gin.New(),
	}

	opts.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
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
