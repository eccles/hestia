package widgets

import (
	"os"

	"github.com/eccles/hestia/pkg/service"
)

//
// Generate Go code from protobuf definition file.
//go:generate ../../../scripts/protoc.sh pkg/services/widgets/widgets.proto

// Generate boilerplate for grpcserver.go and service.go
// First argument value=1 enables GRPC
//go:generate ../../../scripts/grpcserver.sh 1 pkg/services/widgets

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
		cfg: service.Config{
			GRPCServerPort: port(),
		},
	}

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
