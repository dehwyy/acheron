package server

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

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

	r.Any("/:streamName/playlist.m3u8", func(c *gin.Context) {
		opts.Log.Info().Msgf("Request to playlist.m3u8 for <%s>", c.Param("streamName"))

		path := filepath.Join("content", "streams", c.Param("streamName"), "playlist.m3u8")

		playlist, err := os.ReadFile(path)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Data(http.StatusOK, "application/x-mpegurl", playlist)

	})

	r.Any("/:streamName/:segmentName", func(c *gin.Context) {
		opts.Log.Info().Msgf("Request to segment for <%s> and <%s>", c.Param("streamName"), c.Param("segmentName"))

		// TODO: validate url
		path := filepath.Join("content", "streams", c.Param("streamName"), c.Param("segmentName"))

		segment, err := os.ReadFile(path)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Data(http.StatusOK, "video/MP2T", segment)
	})

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
