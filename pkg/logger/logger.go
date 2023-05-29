package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

// Logger represents an interface to a typical logger including logging
// levels, format control and others to be elucidated in future commits.
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

	*slog.Logger
}

// Open creates logger as to the desired loglevel.
func (l *Logger) Open() error {
	var slogger *slog.Logger
	var options slog.HandlerOptions
	if l.Level == "DEBUG" {
		options.AddSource = true
		options.Level = slog.LevelDebug
	} else {
		options.Level = slog.LevelInfo
	}
	if l.Console {
		slogger = slog.New(slog.NewTextHandler(os.Stderr, &options))
	} else {
		// if l.Sampling == 0 {
		//	l.Sampling = defaultSampling
		//}
		slogger = slog.New(slog.NewJSONHandler(os.Stdout, &options))
	}

	if l.ServiceName != "" {
		slogger = slogger.With("service", l.ServiceName)
	}

	l.Logger = slogger

	return nil
}

// Close should be deferred when a Logger is opened.
func (l *Logger) Close() {
	// do nothing for the moment
}
