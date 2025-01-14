package logg

import (
	"slices"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type hookOpts struct {
	sentry *sentry.Client
}

func newHooks(opts hookOpts) []zerolog.Hook {
	return []zerolog.Hook{
		&sentryHook{
			sentryClient: opts.sentry,
		},
	}
}

type sentryHook struct {
	sentryClient *sentry.Client
}

func (h *sentryHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	// It's ok for `sentryClient` to be nil
	if slices.Contains([]zerolog.Level{zerolog.FatalLevel, zerolog.ErrorLevel}, level) && h.sentryClient != nil {
		defer h.sentryClient.Flush(5)
		e.Msg("Sent to Sentry.")

		var sentryLevel sentry.Level
		if level == zerolog.ErrorLevel {
			sentryLevel = sentry.LevelError
		} else {
			sentryLevel = sentry.LevelFatal
		}

		event := h.sentryClient.EventFromMessage(msg, sentryLevel)
		h.sentryClient.CaptureEvent(event, &sentry.EventHint{
			Data: map[string]any{
				"Stack": e.Stack(),
			},
			Context: e.GetCtx(),
		}, nil)
	}
}
