package http

import (
	"log/slog"
	"net/http"

	"github.com/aperezgdev/task-it-api/internal/infrastructure/config"
)

type Server struct {
	slog       *slog.Logger
	httpServer *http.Server
	handler    *http.ServeMux
}

func NewServer(slog *slog.Logger, config config.Config) *Server {
	handler := http.NewServeMux()

	server := http.Server{
		Handler: handler,
		Addr:    ":" + config.ServerPort,
	}

	return &Server{
		slog:       slog,
		httpServer: &server,
		handler:    handler,
	}
}

func (hs *Server) AddHandler(pattern string, handler http.HandlerFunc) {
	hs.handler.HandleFunc(pattern, handler)
}

func (hs *Server) Handler() *http.ServeMux {
	return hs.handler
}

func (hs *Server) Start() error {
	return hs.httpServer.ListenAndServe()
}