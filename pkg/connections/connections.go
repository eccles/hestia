package connections

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eccles/hestia/pkg/logger"
)

type Connections struct{}

func (c *Connections) Connect(logger logger.LoggerInterface) error {
	// put open statements of all other connections to external services
	// here.
	// ....

	// this is where the service stops and pauses for any tern=mination from
	// k8s or other orchestration tools.
	logger.Info().Msg("Wait for termination signal")

	// wait here for termination signal
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
	<-sCh

	// k8s is in charge now so undo handling of signals
	signal.Reset(syscall.SIGINT, syscall.SIGTERM)

	return nil
}
