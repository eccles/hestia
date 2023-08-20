package widgetsService

import (
	"fmt"
	"net"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCService represents an incoming connection i.e a server that handles GRPC messages
// sent from a client.
type GRPCService struct {
	Port string

	Server *grpc.Server
}

func (g *GRPCService) Stop() {
	g.Server.GracefulStop()
}

// StartGRPCService initializes the GRPC service structs generated in thee qpis
// package and starts the GRPC service.
func (s *Service) StartGRPCService() error {
	s.Logger.Info("Start GRPCService %s", s.Version)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	widgetsAPI.RegisterWidgetsServer(grpcServer, s) // s is of type widgetsAPI.WidgetsServer
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":"+s.GRPC.Port)
	if err != nil {
		return fmt.Errorf("listen ':%s' failure: %w", s.GRPC.Port, err)
	}

	s.GRPC.Server = grpcServer
	g := new(errgroup.Group)
	g.Go(func() error {
		err = s.GRPC.Server.Serve(listen)
		if err != nil {
			s.Logger.Info("Failed to start")
			return err
		}
		return nil
	})

	return g.Wait()
}
