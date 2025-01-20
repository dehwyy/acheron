package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func (s *Server) Start(_ context.Context, port uint) error {
	return s.Run(fmt.Sprintf(":%d", port))
}

func (s *Server) Stop(_ context.Context) error {
	return nil
}
