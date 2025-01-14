package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/gqlgen"
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/server"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	Server   *server.Server
	Resolver gqlgen.ResolverRoot
	Log      logg.Logger
}

func NewFx(opts Opts) *handler.Server {
	cfg := gqlgen.Config{
		Resolvers: opts.Resolver,
	}

	schema := gqlgen.NewExecutableSchema(cfg)

	h := handler.New(schema)
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.POST{})

	h.Use(extension.Introspection{})

	opts.Server.Any("/", func(ctx *gin.Context) {
		playground.Handler("GraphQL playground", "/api/query").ServeHTTP(ctx.Writer, ctx.Request)
	})
	opts.Server.Any("/api/query", func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})

	opts.Log.Info().Msg("GraphQL initialized!")

	return h
}
