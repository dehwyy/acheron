package routers

import (
	"github.com/dehwyy/mugen/apps/stream_broadcaster/internal/repos"
	"github.com/dehwyy/mugen/libraries/go/logg"
	"go.uber.org/fx"
)

type PlaylistRouterOpts struct {
	fx.In

	Log      logg.Logger
	FileRepo *repos.FileRepository
}

func NewPlaylistRouterFx(opts PlaylistRouterOpts) *PlaylistRouter {
	return &PlaylistRouter{Log: opts.Log, FileRepo: opts.FileRepo}
}
