package logg

import (
	"io"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"go.uber.org/fx"
)

type Logger = *zerolog.Logger

type Opts struct {
	ServiceName string
}

type Params struct {
	fx.In

	Sentry *sentry.Client `optional:"true"`
}

func New(opts Opts) func(Params) Logger {
	return func(p Params) Logger {
		l := NewFx(p)
		l.Info().Msgf("Logger for <%s> initialized!", opts.ServiceName)
		return l
	}
}

func NewFx(p Params) Logger {
	var out io.Writer

	production := os.Getenv("PRODUCTION") == "TRUE"

	if production {
		out = os.Stdout
	} else {
		// Pretty logging
		out = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC822,
		}
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.ErrorStackFieldName = "stack"

	logger := zerolog.New(out).With().Caller().Timestamp().Logger()

	logger.Hook(newHooks(hookOpts{p.Sentry})...)
	return &logger
}
