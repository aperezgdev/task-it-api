package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"testing"

	domain_errors "github.com/aperezgdev/task-it-api/internal/domain/errors"
)

func TestWriterError(t *testing.T) {
	t.Parallel()
	t.Run("should return internal server error", func(t *testing.T) {
		writter := httptest.NewRecorder()
		writeError(writter, errors.ErrUnsupported)
		if writter.Code != http.StatusInternalServerError {
			t.Errorf("expected %d, got %d", http.StatusInternalServerError, writter.Code)
		}
	})

	t.Run("should return not found error", func(t *testing.T) {
		writter := httptest.NewRecorder()
		writeError(writter, domain_errors.ErrNotExist)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, writter.Code)
		}
	})

	t.Run("should return bad request error", func(t *testing.T) {
		writter := httptest.NewRecorder()
		writeError(writter, domain_errors.NewValidationError("field", "message"))

		if writter.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, writter.Code)
		}
	})

}