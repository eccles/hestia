package widgetsservice

import (
	"fmt"

	widgetsapi "github.com/eccles/hestia/pkg/apis/widgets"
)

// implements the widgetsapi.WidgetsServer interfacw.
type Service struct {
	widgetsapi.UnimplementedWidgetsServer

	// An interface as we may want to mock it out in tests.
	Log Logger

	// A concrete implementation of a GRPC service.
	GRPC GRPCService
}

func (s *Service) Run() error {
	err := s.StartGRPCService()
	if err != nil {
		return fmt.Errorf("grpcservice start failure: %w", err)
	}
	defer s.GRPC.Stop()

	return s.Connect()
}
