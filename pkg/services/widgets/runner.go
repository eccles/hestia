package widgets

import (
	"os"

	"github.com/eccles/hestia/pkg/logger"
)

const (
	serviceName = "widgets"
)

// Generate boilerplate for grpcserver.go and service.go
// First argument value=1 enables GRPC
//go:generate ../../../scripts/grpcserver.sh 1 pkg/services/widgets

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
// The Service strucyt specifies the grpc server code and any other interfaces
// to external services.
func Run() {
	var err error

	s := Service{
		Logger: &logger.Logger{
			Level:       logLevel(),
			ServiceName: serviceName,
		},
		GRPCServerPort: port(),
	}

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
