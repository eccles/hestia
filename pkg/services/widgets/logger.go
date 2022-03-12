package widgetsService

import (
	"github.com/rs/zerolog"
)

// LoggerInterface corresponds to most interfaces that have
// an Open/Close interfaces.
type LoggerInterface interface {
	Open() error
	Close()
	Debug() *zerolog.Event
	Fatal() *zerolog.Event
	Info() *zerolog.Event
}
