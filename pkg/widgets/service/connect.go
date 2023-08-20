package widgetsService

import (
	"github.com/eccles/hestia/pkg/terminate"
)

func (s *Service) Connect() error {
	s.Log.Info("Start")
	// put open statements of all other connections to external services
	// here.
	// ....

	// this is where the service stops and pauses for any termination from
	// k8s or other orchestration tools.
	s.Log.Info("Wait for termination signal")
	terminate.WaitForTermination()

	return nil
}
