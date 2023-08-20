package startup

import (
	"errors"
	"os"

	"github.com/eccles/hestia/pkg/logger"
)

var ErrLogLevelUndefined = errors.New("no loglevel specified")

func getLogLevel() string {
	value, ok := os.LookupEnv("LOGLEVEL")
	if !ok {
		panic(ErrLogLevelUndefined)
	}
	return value
}

type Runner func(string, Logger) error

// Run() is run from the main() method in order to isolate
// code that executes os.Exit(). This enables defers in the run
// function in the second argument.
func Run(serviceName string, run Runner) {
	var exitCode int
	logger.New(getLogLevel())
	log := logger.WithServiceName(serviceName)

	err := run(serviceName, log)
	if err != nil {
		log.Info("Error terminating: %v", err)
		exitCode = 1
	}
	log.Info("Shutting down")
	logger.Close()
	os.Exit(exitCode)
}
