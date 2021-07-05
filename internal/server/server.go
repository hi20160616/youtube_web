package server

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"time"

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
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			PanicLog(e)
		}
	}()
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

func PanicLog(_err error) error {
	filePath := "./PanicLog.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("[" + time.Now().Format(time.RFC3339) + "]--------------------------------------\n")
	write.WriteString(_err.Error() + "\n")
	write.Flush()
	return nil
}
