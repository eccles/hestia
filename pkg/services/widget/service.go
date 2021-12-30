package widget

import (
	"os"
	"os/signal"
	"syscall"
)

type Service struct {
	UnimplementedWidgetsServer
	Cfg Config
}

func (s *Service) Run() {
	// connections to storage, message busses etc go here
	// together with relevant defer statements
	// ...
	// wait here for termination signal
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
	<-sCh

	// k8s is in charge now so undo handling of signals
	signal.Reset(syscall.SIGINT, syscall.SIGTERM)
}
