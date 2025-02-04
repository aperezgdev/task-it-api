package controller

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/application/status"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestStatusControllerPost(t *testing.T) {
	t.Parallel()

	t.Run("should create status", func(t *testing.T) {
		statusRepository := new(repository.MockStatusRepository)
		statusRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepositoryMock := repository.MockBoardRepository{}
		boardRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), statusRepository, &boardRepositoryMock))
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"title":"title","board":"` + uuid.String() + `","nextStatus":["01946ba3-ee73-76e6-83a9-33f87a35d6e9"],"previousStatus":["01946ba3-ee73-76e6-83a9-33f87a35d6e9"]}`)))
		writer := httptest.NewRecorder()
		statusController.PostController(writer, *req)

		if writer.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, writer.Code)
		}
	})
	
	t.Run("should return error on invalid status", func(t *testing.T) {
		statusRepository := new(repository.MockStatusRepository)
		statusRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository := new(repository.MockBoardRepository)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), statusRepository, boardRepository))
		uuid, _ := uuid.NewV7()

		writer := httptest.NewRecorder()
		statusController.PostController(writer, *httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"title":"","board":"` + uuid.String() + `","nextStatus":["01946ba3-ee73-76e6-83a9-33f87a35d6e9"],"previousStatus":["01946ba3-ee73-76e6-83a9-33f87a35d6e9"]}`))))

		if writer.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, writer.Code)
		}
	})
}