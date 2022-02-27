package widgets

// widgets - a GRPC Service that implements the widgetsAPI.WidgetsServer interface.
// This file (runer.go must be defined manually) .
//
// Caveats - the packagename and the directory must have the same name - in this
//           case 'widgets'.
//
// Additional files to be created manually are handlers.go and connetions.go.
// All other fioes are generated using 'go generate' commands

import (
	"os"

	"github.com/eccles/hestia/pkg/logger"
)

const (
	serviceName = "widgets"
)

// Generate boilerplate for grpcserver.go and service.go
// First argument value=1 enables GRPC
//go:generate ../../../scripts/grpcserver.sh 1

func logLevel() string {
	val, ok := os.LookupEnv("LOGLEVEL")
	if !ok {
		val = "INFO"
	}

	return val
}

func port() string {
	val, ok := os.LookupEnv("GRPC_SERVICE_PORT")
	if !ok {
		val = ""
	}

	return val
}

// Run initializes the Service struct and executes its Run method.
// The Service struct specifies the grpc server code and any other interfaces
// to external services defined in connections.go.
func Run() {
	var err error

	s := Service{
		Logger: &logger.Logger{
			Level:       logLevel(),
			ServiceName: serviceName,
		},
		GRPC: GRPCService{
			Port: port(),
		},
	}

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
