package logger

// Creates global logger from which all loggers are derived.
import (
	"os"

	"golang.org/x/exp/slog"
)

const (
	serviceLogKey = "service"
)

var RootLogger *slog.Logger //nolint:gochecknoglobals // all loggers derive from this

func New(level string) {
	var options slog.HandlerOptions
	if level == "DEBUG" {
		options.AddSource = true
		options.Level = slog.LevelDebug
		RootLogger = slog.New(slog.NewTextHandler(os.Stderr, &options))
	} else {
		options.Level = slog.LevelInfo
		RootLogger = slog.New(slog.NewJSONHandler(os.Stdout, &options))
	}
}

// close is a dummy function at the moment - it should be deferred.
func Close() {
}

func WithServiceName(serviceName string) Logger {
	return RootLogger.With(serviceLogKey, serviceName)
}
