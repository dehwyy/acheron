package types

type DefaultRequest[T Payload] struct {
	payload T
}

func NewRequest[T Payload](p T) Request[T] {
	return &DefaultRequest[T]{payload: p}
}

func (r DefaultRequest[T]) Get() T {
	return r.payload
}
