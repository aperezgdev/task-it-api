package controller

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/application/user"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

func TestUserControllerPost(t *testing.T) {
	t.Parallel()

	t.Run("should create user", func(t *testing.T) {
		writter := httptest.NewRecorder()
		userRepositoryMock := repository.MockUserRepository{}
		userRepositoryMock.On("Save", mock.Anything, mock.Anything).Return(nil)
		userController := NewUserController(*slog.Default(), user.NewUserCreator(*slog.Default(), &userRepositoryMock))
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"email": "email@email.com"}`))
		userController.PostController(writter, req)

		if writter.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, writter.Code)
		}
	})

	t.Run("should return bad request", func(t *testing.T) {
		writter := httptest.NewRecorder()
		userController := NewUserController(*slog.Default(), user.NewUserCreator(*slog.Default(), nil))
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"email": ""}`))
		userController.PostController(writter, req)

		if writter.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, writter.Code)
		}
	})
}