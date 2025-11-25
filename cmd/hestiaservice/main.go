package main

import (
	"fmt"
	"os"

	"github.com/eccles/hestia/httpserver"
	"github.com/eccles/hestia/logger"
	"github.com/eccles/hestia/services/hestia"
	"github.com/eccles/hestia/startup"
)

const (
	serviceName = "hestiat"
)

func main() {
	startup.Run(serviceName, run)
}

func run(log logger.Logger) error {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		err := fmt.Errorf("required environment variable is not defined: %s", "PORT")
		log.Info(err.Error())
		return err
	}

	service := hestia.New(
		serviceName,
		log,
		&hestia.Config{},
	)
	defer service.Close()

	h := httpserver.New(
		log,
		serviceName,
		port,
		service.Mux(),
	)

	s := startup.NewListeners(
		log,
		serviceName,
		startup.WithListeners(h),
	)
	return s.Listen() // blocks until either one listener fails or sigterm is received.
}
