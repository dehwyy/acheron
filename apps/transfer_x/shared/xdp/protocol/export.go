package xdp

import t "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol/types"

type Payload = t.Payload
type StreamPayload = t.StreamPayload

type Request[T Payload] = t.Request[T]
type Response[T Payload] = t.Response[T]
type StreamRequest[T StreamPayload] = t.StreamRequest[T]
type StreamResponse[T StreamPayload] = t.StreamResponse[T]
