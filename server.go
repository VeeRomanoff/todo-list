package todo_app

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *logrus.Logger
}

func NewServer() *Server {
	return &Server{
		logger: logrus.New(),
	}
}

func (s *Server) RunServer(port string, handler http.Handler) error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	s.logger.Info("server started on port ", port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) StopServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
