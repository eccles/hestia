package widgets

import (
	"os"
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

func Run() {
	var err error

	s := Service{
		LogLevel:       logLevel(),
		GRPCServerPort: port(),
	}

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
