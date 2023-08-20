package logger

import (
	"golang.org/x/exp/slog"
)

// Logger corresponds to most logger interfaces.
type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)

	With(...any) *slog.Logger
}
