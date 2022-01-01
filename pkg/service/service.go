package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type GRPCServer interface {
	Start() error
	GracefulStop()
}

type Service struct {
	GRPCServer GRPCServer
}

func (s *Service) Run() error {
	if s.GRPCServer != nil {
		err := s.GRPCServer.Start()
		if err != nil {
			return fmt.Errorf("start failure: %w", err)
		}
		defer s.GRPCServer.GracefulStop()
	}

	// wait here for termination signal
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
	<-sCh

	// k8s is in charge now so undo handling of signals
	signal.Reset(syscall.SIGINT, syscall.SIGTERM)

	return nil
}
