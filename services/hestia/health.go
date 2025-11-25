package hestia

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// NotServing changes status to NOT_SERVING.
func (s *Service) NotServing() {
	// s.Lock()
	// defer s.Unlock()
	s.HealthStatus = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	s.log.Info("Health set to 'NOT_SERVING'")
}

// Health implements health check.
func (s *Service) Health(w http.ResponseWriter, r *http.Request) {
	// s.RLock()
	// defer s.RUnlock()
	if s.HealthStatus == grpc_health_v1.HealthCheckResponse_SERVING {
		w.WriteHeader(200)
		fmt.Fprint(w, "OK")
		return
	}
	s.log.Debug("Health check: 'NOT_SERVING'")
	w.WriteHeader(500)
	fmt.Fprint(w, "NOT_SERVING")
}
