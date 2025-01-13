package gqlresolvers

import (
	"context"
	"fmt"

	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/gqlgen"
	gqlmodels "github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/models"
)

func (r *queryResolver) GetStream(ctx context.Context, req gqlmodels.StreamRequest) (*gqlmodels.StreamResponse, error) {
	panic(fmt.Errorf("not implemented: GetStream - getStream"))
}

func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
