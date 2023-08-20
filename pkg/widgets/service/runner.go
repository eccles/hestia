package widgetsService

// widgets - a GRPC Service that implements the widgetsAPI.WidgetsServer interface.
// This file (runner.go must be defined manually) .
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
		Log: logger.WithServiceName(serviceName),
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
