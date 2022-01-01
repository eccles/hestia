package widget

import (
	"os"

	"github.com/eccles/hestia/pkg/service"
)

func port() string {
	val, ok := os.LookupEnv("WIDGETS_GRPC_SERVICE_PORT")
	if !ok {
		val = ""
	}

	return val
}

func Run() {
	var err error

	gServer := grpcService{
		Port: port(),
	}
	s := service.Service{
		GRPCServer: &gServer,
	}

	err = s.Run()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
