package widget

import (
	"net"
)

type GRPCServer interface {
	Serve(net.Listener) error
	GracefulStop()
}

type Config struct {
	GRPCServer GRPCServer
}
