package interceptors

import (
	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/intervalpli"
)

func New() (*interceptor.Registry, error) {
	registry := &interceptor.Registry{}

	intervalPliFactory, err := intervalpli.NewReceiverInterceptor()
	if err != nil {
		return nil, err
	}

	registry.Add(intervalPliFactory)

	return registry, nil
}
