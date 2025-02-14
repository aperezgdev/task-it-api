package controller

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthController(t *testing.T) {
	t.Parallel()

	t.Run("should return OK on valid request", func(t *testing.T) {
		healthController := NewHealthController(*slog.Default())
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/health", nil)
		healthController.GetHealth(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("Response code should be 200")
		}
	})
}