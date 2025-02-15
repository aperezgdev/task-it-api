package controller

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/application/board"
	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestBoardControllerPost(t *testing.T) {
	t.Parallel()

	t.Run("should create board", func(t *testing.T) {
		writter := httptest.NewRecorder()
		boardRepositoryMock := repository.MockBoardRepository{}
		boardRepositoryMock.On("Save", mock.Anything, mock.Anything).Return(nil)
		userRepositoryMock := repository.MockUserRepository{}
		userRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		teamRepositoryMock := repository.MockTeamRepository{}
		teamRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		boardController := NewBoardController(*slog.Default(), board.NewBoardCreator(*slog.Default(), &boardRepositoryMock, &userRepositoryMock, &teamRepositoryMock), board.NewBoardRemover(*slog.Default(), nil), board.NewBoardFinderByTeam(*slog.Default(), &boardRepositoryMock, &teamRepositoryMock))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"title": "title", "description": "description", "owner": "` + uuid.String() + `", "team": "` + uuid.String() + `"}`))
		boardController.PostController(writter, req)

		if writter.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, writter.Code)
		}
	})

	t.Run("should return bad request", func(t *testing.T) {
		writter := httptest.NewRecorder()
		boardController := NewBoardController(*slog.Default(), board.NewBoardCreator(*slog.Default(), nil, nil, nil), board.NewBoardRemover(*slog.Default(), nil), board.NewBoardFinderByTeam(*slog.Default(), nil, nil))
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"title": ""}`))
		boardController.PostController(writter, req)

		if writter.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, writter.Code)
		}
	})
}

func TestBoardControllerDelete(t *testing.T) {
	t.Parallel()

	t.Run("should delete board", func(t *testing.T) {
		writter := httptest.NewRecorder()
		boardRepositoryMock := repository.MockBoardRepository{}
		boardRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		boardRepositoryMock.On("Delete", mock.Anything, mock.Anything).Return(nil)
		boardController := NewBoardController(*slog.Default(), board.NewBoardCreator(*slog.Default(), nil, nil, nil), board.NewBoardRemover(*slog.Default(), &boardRepositoryMock), board.NewBoardFinderByTeam(*slog.Default(), nil, nil))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodDelete, "/"+uuid.String(), nil)
		req.SetPathValue("id", uuid.String())
		boardController.DeleteController(writter, *req)

		if writter.Code != http.StatusNoContent {
			t.Errorf("expected %d, got %d", http.StatusNoContent, writter.Code)
		}
	})

	t.Run("should return not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		boardRepositoryMock := &repository.MockBoardRepository{}
		boardRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), nil)
		boardController := NewBoardController(*slog.Default(), board.NewBoardCreator(*slog.Default(), boardRepositoryMock, nil, nil), board.NewBoardRemover(*slog.Default(), boardRepositoryMock), board.NewBoardFinderByTeam(*slog.Default(), nil, nil))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodDelete, "/"+uuid.String(), nil)
		req.SetPathValue("id", uuid.String())
		boardController.DeleteController(writter, *req)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})
}

func TestBoardControllerGetControllerByTeam(t *testing.T) {
	t.Parallel()

	t.Run("should return board by team", func(t *testing.T) {
		writter := httptest.NewRecorder()
		boardRepositoryMock := repository.MockBoardRepository{}
		teamRepositoryMock := repository.MockTeamRepository{}
		teamRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		boardRepositoryMock.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.NewOptional([]model.Board{{}}), nil)
		boardController := NewBoardController(*slog.Default(), board.NewBoardCreator(*slog.Default(), nil, nil, nil), board.NewBoardRemover(*slog.Default(), nil), board.NewBoardFinderByTeam(*slog.Default(), &boardRepositoryMock, &teamRepositoryMock))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodGet, "/"+uuid.String(), nil)
		req.SetPathValue("teamId", uuid.String())
		boardController.GetControllerByTeam(writter, *req)

		if writter.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, writter.Code)
		}
	})

	t.Run("should return not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		boardRepositoryMock := repository.MockBoardRepository{}
		teamRepositoryMock := repository.MockTeamRepository{}
		teamRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), errors.ErrNotExist)
		boardRepositoryMock.On("FindByTeam", mock.Anything, mock.Anything).Return([]model.Board{}, nil)
		boardController := NewBoardController(*slog.Default(), board.NewBoardCreator(*slog.Default(), nil, nil, nil), board.NewBoardRemover(*slog.Default(), nil), board.NewBoardFinderByTeam(*slog.Default(), &boardRepositoryMock, &teamRepositoryMock))
		uuid, _ := uuid.NewV7()
		
		req := httptest.NewRequest(http.MethodGet, "/"+uuid.String(), nil)
		req.SetPathValue("teamId", uuid.String())
		boardController.GetControllerByTeam(writter, *req)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})
}