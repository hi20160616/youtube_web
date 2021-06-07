package server

import (
	"context"
	"net/http"

	"github.com/hi20160616/youtube_web/internal/server/handler"
)

type Server struct {
	http.Server
}

func NewServer(address string) (*Server, error) {
	return &Server{http.Server{
		Addr:    address,
		Handler: handler.GetHandler(),
	}}, nil
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return ctx.Err()
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.Shutdown(context.Background()); err != nil {
		return err
	}
	return ctx.Err()
}
