package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

const (
	defaultSampling = 10
)

//go:generate mockery --all --dry-run=false

// LoggerInterface corresponds to most interfaces that have
// an Open/Close interfaces.
type LoggerInterface interface {
	Open() error
	Close()
	Info() *zerolog.Event
	Debug() *zerolog.Event
}

// Logger represents an interface to a typical logger including logging
// levels, format control and others to be elucidated in fiture commits.
type Logger struct {

	// Service - the logger has this field add as "service"
	ServiceName string

	// Level specifies the logging level. "DEBUG" is debugging, anything else is INFO
	Level string

	// Console
	//
	// If true then console style output is used - usually for client-side applications.
	// If false log output contains timestamp and a sampling interface.
	Console bool

	// Sampling - when console is false then log output is sampled to mitigate sudden
	// bursts of output typically when an error of some kind occurs.
	Sampling uint32

	zerolog.Logger
}

// Open creates logger as to the desired loglevel.
func (l *Logger) Open() error {
	var zlogger zerolog.Logger
	if l.Console {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

		zlogger = zerolog.New(output).
			With().
			Timestamp().
			Logger()
	} else {
		if l.Sampling == 0 {
			l.Sampling = defaultSampling
		}
		zlogger = zerolog.New(os.Stderr).
			With().
			Timestamp().
			Logger()
		zlogger = zlogger.Sample(&zerolog.BasicSampler{N: l.Sampling})
	}

	if l.ServiceName != "" {
		zlogger = zlogger.With().Str("service", l.ServiceName).Logger()
	}
	if l.Level == "DEBUG" {
		zlogger.Level(zerolog.DebugLevel)
	} else {
		zlogger.Level(zerolog.InfoLevel)
	}

	l.Logger = zlogger

	return nil
}

// Close should be deferred when a Logger is opened.
func (l *Logger) Close() {
	// do nothing for the moment
}
