package service

import (
	"context"
	"log"
	"net/http"

	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"github.com/hi20160616/youtube_web/internal/pkg/handler"
)

type Server struct {
	http.Server
}

func NewServer(address string) (*Server, error) {
	h, err := handler.NewHandler(db.YoutubeService)
	if err != nil {
		return nil, err
	}
	return &Server{http.Server{
		Addr:    address,
		Handler: h.GetHandler(),
	}}, nil
}

func (s *Server) Start(ctx context.Context) error {
	log.Println("Server start on " + s.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return ctx.Err()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Server stop ...")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
	log.Println("Server gracefully stopped.")
	return ctx.Err()
}
