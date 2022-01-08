package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

const (
	sampling = 10
)

type Options struct {
	console bool
}

type ResourceOption func(*Options)

func WithConsole() ResourceOption {
	return func(r *Options) {
		r.console = true
	}
}

type Logger struct {
	zerolog.Logger
}

// New creates logger as to the desired loglevel.
func New(level string, opts ...interface{}) Logger {
	options := Options{}
	for _, opt := range opts {
		if fopt, ok := opt.(ResourceOption); ok {
			fopt(&options)
		}
	}

	var zlogger zerolog.Logger
	if options.console {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

		zlogger = zerolog.New(output).With().Timestamp().Logger()
	} else {
		zlogger = zerolog.New(os.Stderr).
			With().
			Timestamp().
			Logger()
		zlogger = zlogger.Sample(&zerolog.BasicSampler{N: sampling})
	}

	if level == "DEBUG" {
		zlogger.Level(zerolog.DebugLevel)
	} else {
		zlogger.Level(zerolog.InfoLevel)
	}

	return Logger{zlogger}
}
