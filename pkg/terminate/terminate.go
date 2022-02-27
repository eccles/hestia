package terminate

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForTermination() {
	// wait here for termination signal
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
	<-sCh

	// k8s is in charge now so undo handling of signals
	signal.Reset(syscall.SIGINT, syscall.SIGTERM)
}
