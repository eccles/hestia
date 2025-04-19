package widgetsservice

// widgets - a GRPC Service that implements the widgetsapi.WidgetsServer interface.
// This file (runner.go must be defined manually) .
//
// Additional files to be created manually are handlers.go and connetions.go.
// All other fioes are generated using 'go generate' commands

import (
	"fmt"
	"os"
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
// to external services defined in connect.go.
func Run(serviceName string, log Logger) error {
	var err error

	s := Service{
		Log: log,
		GRPC: GRPCService{
			Port: port(),
		},
	}

	err = s.Run()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
