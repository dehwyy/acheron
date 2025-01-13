package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct {
	*gin.Engine
}

func (s *Server) Start(_ context.Context, port uint) error {
	return s.Run(fmt.Sprintf(":%d", port))
}

func (*Server) Stop(_ context.Context) error {
	return nil
}

type Opts struct {
	fx.In
	LC fx.Lifecycle
}

func New(opts Opts) *Server {
	r := &Server{
		gin.New(),
	}

	opts.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// TODO: Logging
				r.Start(ctx, 8081)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return r.Stop(ctx)
		},
	})

	return r
}
