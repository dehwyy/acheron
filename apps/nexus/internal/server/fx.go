package server

import (
	"context"

	"github.com/dehwyy/acheron/libraries/go/config"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"github.com/gin-contrib/cors"
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

	opts.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				opts.Log.Info().Msg("Starting server...")
				opts.Log.Fatal().Msgf("%v", r.Start(ctx, opts.Config.Addr().Ports.Nexus))
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

// func serveFile(c *gin.Context, filepath string) {
// 	// TODO: validate url
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		c.Error(err)
// 		return
// 	}

// 	io.Copy(c.Writer, file)
// 	c.Writer.Flush()

// 	// TODO: package `file`
// }
