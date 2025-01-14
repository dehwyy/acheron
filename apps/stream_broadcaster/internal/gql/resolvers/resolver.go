package gqlresolvers

import "github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/gqlgen"

type Resolver struct{}

func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }
