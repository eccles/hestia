package logger

import (
	"log/slog"
)

// Logger corresponds to most logger interfaces.
type Logger interface {
	Debug(fmt string, vals ...any)
	Info(fmt string, vals ...any)

	With(vals ...any) *slog.Logger
}
