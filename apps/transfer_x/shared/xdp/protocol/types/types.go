package types

type Payload interface{}
type StreamPayload interface{}

type Request[P Payload] interface {
	Get() P
}
type Response[P Payload] interface {
	Get() P
}
type StreamRequest[P Payload] interface {
	Get() P
}
type StreamResponse[P Payload] interface {
	Get() P
}
