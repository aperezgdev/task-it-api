package controller

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/application/status"
	"github.com/aperezgdev/task-it-api/internal/domain/errors"
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
		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), statusRepository, &boardRepositoryMock), status.NewStatusRemover(*slog.Default(), nil), status.StatusFinderByBoard{})
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
		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), statusRepository, boardRepository), status.NewStatusRemover(*slog.Default(), nil), status.StatusFinderByBoard{})
		uuid, _ := uuid.NewV7()

		writer := httptest.NewRecorder()
		statusController.PostController(writer, *httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"title":"","board":"` + uuid.String() + `","nextStatus":["01946ba3-ee73-76e6-83a9-33f87a35d6e9"],"previousStatus":["01946ba3-ee73-76e6-83a9-33f87a35d6e9"]}`))))

		if writer.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, writer.Code)
		}
	})
}

func TestStatusControllerDelete(t *testing.T) {
	t.Parallel()

	t.Run("should delete status", func(t *testing.T) {
		statusRepository := new(repository.MockStatusRepository)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)
		statusRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), nil, nil), status.NewStatusRemover(*slog.Default(), statusRepository), status.StatusFinderByBoard{})
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodDelete, "/"+uuid.String(), nil)
		req.SetPathValue("id", uuid.String())
		writer := httptest.NewRecorder()
		statusController.DeleteController(writer, *req)

		if writer.Code != http.StatusNoContent {
			t.Errorf("expected %d, got %d", http.StatusNoContent, writer.Code)
		}
	})

	t.Run("should return not found", func(t *testing.T) {
		statusRepository := new(repository.MockStatusRepository)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Status](), nil)
		statusRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), nil, nil), status.NewStatusRemover(*slog.Default(), statusRepository), status.StatusFinderByBoard{})
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodDelete, "/"+uuid.String(), nil)
		req.SetPathValue("id", uuid.String())
		writer := httptest.NewRecorder()
		statusController.DeleteController(writer, *req)

		if writer.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writer.Code)
		}
	})
}

func TestStatusControllerGetControllerByBoard(t *testing.T) {
	t.Parallel()

	t.Run("should return status by board", func(t *testing.T) {
		writter := httptest.NewRecorder()
		statusRepositoryMock := repository.MockStatusRepository{}
		boardRepositoryMock := repository.MockBoardRepository{}
		boardRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		statusRepositoryMock.On("FindByBoard", mock.Anything, mock.Anything).Return(pkg.NewOptional([]model.Status{}), nil)

		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), &statusRepositoryMock, &boardRepositoryMock), status.NewStatusRemover(*slog.Default(), nil), status.NewStatusFinderByBoard(*slog.Default(), &statusRepositoryMock, &boardRepositoryMock))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodGet, "/"+uuid.String(), nil)
		req.SetPathValue("boardId", uuid.String())
		statusController.GetControllerByBoard(writter, *req)

		if writter.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, writter.Code)
		}
	})

	t.Run("should return not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		statusRepositoryMock := repository.MockStatusRepository{}	
		boardRepositoryMock := repository.MockBoardRepository{}	
		boardRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), errors.ErrNotExist)
		statusRepositoryMock.On("FindByBoard", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[[]model.Status](), errors.ErrNotExist)

		statusController := NewStatusController(*slog.Default(), status.NewStatusCreator(*slog.Default(), &statusRepositoryMock, &boardRepositoryMock), status.NewStatusRemover(*slog.Default(), nil), status.NewStatusFinderByBoard(*slog.Default(), &statusRepositoryMock, &boardRepositoryMock))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodGet, "/"+uuid.String(), nil)
		req.SetPathValue("boardId", uuid.String())
		statusController.GetControllerByBoard(writter, *req)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})
}