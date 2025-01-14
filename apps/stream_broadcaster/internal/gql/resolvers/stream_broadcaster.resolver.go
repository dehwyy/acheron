package gqlresolvers

import (
	"context"
	"fmt"

	gqlmodels "github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetStream(ctx context.Context, req gqlmodels.StreamRequest) (*gqlmodels.StreamResponse, error) {
	panic(fmt.Errorf("not implemented: GetStream - getStream"))
}
