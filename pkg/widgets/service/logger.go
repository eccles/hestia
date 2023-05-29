package widgetsService

// LoggerInterface corresponds to most logger interfaces.
type LoggerInterface interface {
	Open() error
	Debug(string, ...any)
	Info(string, ...any)
}
