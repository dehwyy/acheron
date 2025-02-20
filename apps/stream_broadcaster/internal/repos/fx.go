package repos

import (
	"github.com/dehwyy/acheron/libraries/go/config"
	"github.com/dehwyy/acheron/libraries/go/logg"
	"go.uber.org/fx"
)

type FileRepositoryOpts struct {
	fx.In

	Log    logg.Logger
	Config config.Config
}

func NewFileRepoFx(opts FileRepositoryOpts) *FileRepository {
	return (&FileRepository{
		Log:                  opts.Log,
		M3u8StreamsDirectory: opts.Config.M3u8().StreamsDirectory,
	}).prepare()
}
