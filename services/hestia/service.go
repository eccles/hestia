package hestia

import (
	"net/http"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type HTTPHandlerFunc = func(http.Handler) http.Handler

// Service implements handlers.
type Service struct {

	// make a copy so that users cannot change anything after the service has started
	cfg Config

	HealthStatus grpc_health_v1.HealthCheckResponse_ServingStatus

	log Logger

	name string
}

// New creates a new hestia service instance.
func New(name string, log Logger, cfg *Config) Service {
	return Service{
		name:         "httpmux" + name,
		cfg:          *cfg,
		log:          log.With("httpmux", name),
		HealthStatus: grpc_health_v1.HealthCheckResponse_SERVING,
	}
}

func (s *Service) String() string {
	return s.name
}

// Open will instantiate any inter-service communication channels.
func (s *Service) Open() error {
	return nil
}

// Close will close any inter-service communication channels.
func (s *Service) Close() {
}
