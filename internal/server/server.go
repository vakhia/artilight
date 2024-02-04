package server

import (
	"github.com/vakhia/artilight/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpRouter *http.Server
}

func NewServer(config *config.Config, handler http.Handler) *Server {
	return &Server{
		httpRouter: &http.Server{
			Addr:           ":" + config.Port,
			Handler:        handler,
			ReadTimeout:    18000 * time.Second,
			WriteTimeout:   18000 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpRouter.ListenAndServe()
}
