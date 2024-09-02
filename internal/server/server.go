package server

import (
	"context"
	"github.com/fanfaronDo/to_do/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HttpServer, httpHandler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		Handler:        httpHandler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
