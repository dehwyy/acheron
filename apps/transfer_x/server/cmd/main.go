package main

import (
	"crypto/tls"

	xdp "github.com/dehwyy/acheron/apps/transfer_x/shared/xdp/protocol"
)

func main() {
	srv, err := xdp.NewXDPServer(xdp.ServerParams{
		TLS: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	})
	if err != nil {
		panic(err)
	}

	r := xdp.NewRouter()
	xdp.AddRoute(r, "/", func(req xdp.Request[P]) error {
		return nil
	})

	xdp.AddStreamingRoute(r, "/", func(rx <-chan xdp.StreamRequest[P], tx chan<- xdp.StreamResponse[xdp.StreamPayload]) error {
		for req := range rx {
			_ = req.Get()
			tx <- P{1, 2}
		}
		return nil
	})

	defer srv.Stop()
	if err = srv.Start(r); err != nil {
		panic(err)
	}
}

type SomeEndpoint struct{}

type P struct {
	Value int  `xd:"value"`
	Alo   uint `xd:"dd"`
}

func (p P) Get() xdp.StreamPayload {
	return p
}

// func R() {
// 	p := P{1, 2}
// 	t := reflect2.TypeOf(p)

// 	for i := 0; i < t.Type1().NumField(); i++ {
// 		field := t.Type1().Field(i)
// 		fmt.Printf("Field: %s, Type: %s, Tag: %s\n", field.Name, field.Type, field.Tag.Get("xd"))
// 	}
// }
