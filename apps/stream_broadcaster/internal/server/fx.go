package server

import (
	"context"

	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	LC  fx.Lifecycle
	Log logg.Logger
}

func NewFx(opts Opts) *Server {
	r := &Server{
		gin.New(),
	}

	opts.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				opts.Log.Info().Msg("Starting server...")
				// TODO: add config
				r.Start(ctx, 8081)
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
