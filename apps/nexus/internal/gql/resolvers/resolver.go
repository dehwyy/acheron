package resolvers

import (
	"github.com/dehwyy/acheron/libraries/go/logg"
	"go.uber.org/fx"
)

type Deps struct {
	fx.In

	Log logg.Logger
}

type Resolver struct {
	Deps
}

func New(deps Deps) *Resolver {
	return &Resolver{deps}
}
