package widget

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcService struct {
	UnimplementedWidgetsServer

	grpcServer *grpc.Server
	Port       string
}

func (g *grpcService) Start() error {
	g.grpcServer = grpc.NewServer()

	RegisterWidgetsServer(g.grpcServer, g)
	reflection.Register(g.grpcServer)

	listen, err := net.Listen("tcp", ":"+g.Port)
	if err != nil {
		return fmt.Errorf("listen ':%s' failure: %w", g.Port, err)
	}

	go func() {
		err = g.grpcServer.Serve(listen)
		if err != nil {
			panic("Failed to start")
		}
	}()

	return nil
}

func (g *grpcService) GracefulStop() {
	g.grpcServer.GracefulStop()
}
