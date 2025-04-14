package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/marcopaulosilva/poc_devin/internal/infrastructure/logger"
)

type Server struct {
	server *http.Server
	logger logger.Logger
}

func NewServer(port int, handler http.Handler, logger logger.Logger) *Server {
	return &Server{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      handler,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Info("Starting API server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down API server")
	return s.server.Shutdown(ctx)
}
