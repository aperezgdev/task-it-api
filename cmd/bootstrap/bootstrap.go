package bootstrap

import (
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/infrastructure/config"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/http"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/http/controller"
)

func Run() error {
	logger := slog.Default()
	config := config.NewConfig(logger)
	server := http.NewServer(logger, config)

	server.AddHandler("/health", controller.NewHealthController(*logger).GetHealth)

	return server.Start()
}