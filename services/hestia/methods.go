package hestia

import (
	"net/http"
)

// Method is a simple method that can be connected to an endpoint in the mux.
// Nothing implemented yet.
func (s *Service) Method(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
}
