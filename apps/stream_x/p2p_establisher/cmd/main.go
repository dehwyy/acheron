package main

import (
	"github.com/dehwyy/acheron/apps/stream_x/p2p_establisher/internal/server"
	"github.com/dehwyy/acheron/apps/stream_x/p2p_establisher/internal/server/routers"
	"github.com/dehwyy/acheron/libraries/go/config"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"go.uber.org/fx"
)

// @title WhipWhep API
// @version 1.0
// @description **Whip** and **Whep** protocol implementation. `OBS => Server` - *Whip*. `Server <=> Client` (e.g from browser) - *Whep*. It's a **bridge** to make `WebRTC` connection.

// @tag.name WhipWhep
// @tag.description Essential router

// @contact.name dehwyy
// @contact.url	https://t.me/dehwyy
// @contact.email	dehwyyy@gmail.icom

// @externalDocs.description IETF WHIP
// @externalDocs.url https://datatracker.ietf.org/doc/html/draft-ietf-wish-whip-16
// // @externalDocs.description IETF WHEP
// // @externalDocs.url https://datatracker.ietf.org/doc/html/draft-murillo-whep-03

// @host localhost:8080
// @BasePath /api/v1

func main() {
	fx.New(
		fx.Provide(
			config.New(config.ConfigConstructorParams{}),
			logg.New(logg.Opts{
				ServiceName: "stream_whip-whep",
			}),
		),
		fx.Provide(routers.NewWhipWhepRouterFx),
		fx.Invoke(server.NewFx),
	).Run()
}
