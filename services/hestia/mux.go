package hestia

import (
	"net/http"
)

func (s *Service) Mux() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("GET /health", s.Health)

	return m
}
