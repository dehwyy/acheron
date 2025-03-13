package middleware

import (
	"sync/atomic"

	"github.com/dehwyy/acheron/libraries/go/logg"
	"github.com/gin-gonic/gin"
)

func NewLoggerMiddleware(log logg.Logger) func(*gin.Context) {
	var id atomic.Uint64
	return func(ctx *gin.Context) {
		reqID := id.Add(1)
		log.Debug().Msgf("Request ID: %d", reqID)
		log.Debug().Msgf("To <%s> from <%s>", ctx.Request.URL, ctx.ClientIP())

		log.Debug().Msgf("Start processing request ID=%d", reqID)
		ctx.Next()
		log.Debug().Msgf("Finish processing request ID=%d", reqID)

		for _, err := range ctx.Errors {
			log.Error().Msgf(`{"error": %v, "meta": %v}`, err.Err, err.Meta)
		}
	}
}
