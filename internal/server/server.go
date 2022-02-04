package server

import (
	"context"
	"fmt"
	"github.com/alexm24/cache-img/internal/models"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg models.HTTPServerConfig, handler http.Handler) error {
	fmt.Println(cfg)
	s.httpServer = &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
