package controller

import (
	"log/slog"
	"net/http"
)

type HealthController struct {
	logger slog.Logger
}

func NewHealthController(logger slog.Logger) HealthController {
	return HealthController{
		logger,
	}
}

func (h HealthController) GetHealth(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HealthController - GetHealth - Request")
	_, err := w.Write([]byte("OK"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}