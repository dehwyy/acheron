package server

import (
	"context"

	"github.com/dehwyy/mugen/apps/stream_whip/internal/server/middleware"
	"github.com/dehwyy/mugen/apps/stream_whip/internal/server/routers"
	"github.com/dehwyy/mugen/libraries/go/config"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	// Swagger
	_ "github.com/dehwyy/mugen/apps/stream_whip/docs" // import docs
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

type ServerParams struct {
	fx.In

	LC     fx.Lifecycle
	Config config.Config
	Log    logg.Logger

	WhipWhepRouter *routers.WhipWhepRouter
}

func NewFx(params ServerParams) *Server {
	r := &Server{
		gin.New(),
	}

	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

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
	// TODO: remove
	r.StaticFile("/", "./apps/stream_whip/cmd/index.html")

	v1 := r.Group("/api/v1")
	v1.Use(middleware.NewLoggerMiddleware(params.Log))

	params.WhipWhepRouter.RegisterRoutes(v1)

	params.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				params.Log.Fatal().Msgf("%v", r.Start(ctx, params.Config.Addr().Ports.StreamWhip))
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return r.Stop(ctx)
		},
	})

	return r
}
