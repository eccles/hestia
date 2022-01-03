package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	console = "console"
)

// Resource - output to file or console.
type Resource struct {
	console  bool
	filename string
}

type ResourceOption func(*Resource)

func WithFile(filename string) ResourceOption {
	return func(r *Resource) {
		r.filename = filename
	}
}

func WithConsole() ResourceOption {
	return func(r *Resource) {
		r.console = true
	}
}

type Logger struct {
	*zap.Logger
	undoLogger func()
}

// OnExit should be deferred immediately after calling the
// New() method.
func (logger *Logger) OnExit() {
	_ = logger.Sync()
	logger.undoLogger()
}

// New creates logger as to the desired loglevel
// Additionally log output from other loggers in 3rd-party packages
// is redirected to the INFO label of these loggers.
// Both ResourceOption and zap.Option types are supported option types. The
// zap.Options are passed on the to zap logger.
func New(level string, opts ...interface{}) Logger {
	r := &Resource{}

	var zopts []zap.Option
	for _, opt := range opts {
		if fopt, ok := opt.(ResourceOption); ok {
			fopt(r)
		}
		if iopt, ok := opt.(zap.Option); ok {
			zopts = append(zopts, iopt)
		}
	}

	logger := Logger{}
	if level == "NOOP" {
		logger.Logger = zap.NewNop()
		logger.undoLogger = zap.RedirectStdLog(logger.Logger)

		return logger
	}

	var err error
	cfg := zap.NewProductionConfig()
	if r.filename != "" {
		cfg.OutputPaths = []string{r.filename}
	}
	if r.console {
		cfg.Encoding = console
		cfg.EncoderConfig = zapcore.EncoderConfig{
			MessageKey: "message",
		}
	}

	switch level {
	case "DEBUG":
		cfg.Level.SetLevel(zapcore.DebugLevel)
	default:
		cfg.Level.SetLevel(zapcore.InfoLevel)
	}

	logger.Logger, err = cfg.Build(zopts...)
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}

	logger.undoLogger = zap.RedirectStdLog(logger.Logger)

	return logger
}
