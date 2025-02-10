package status

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestStatusFinderByBoard(t *testing.T) {
	t.Run("should find status by board", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		statusRepository.On("FindByBoard", mock.Anything, mock.Anything).Return(pkg.NewOptional([]model.Status{}), nil)

		statusFinder := NewStatusFinderByBoard(*slog.Default(), statusRepository, boardRepository)
		statuses, err := statusFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}		
		if len(statuses) != 0 {
			t.Errorf("Expected empty slice, got %v", statuses)
		}
	})

	t.Run("should return error trying to find status by board", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), errors.ErrNotExist)
		statusRepository.On("FindByBoard", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[[]model.Status](), errors.ErrNotExist)

		statusFinder := NewStatusFinderByBoard(*slog.Default(), statusRepository, boardRepository)
		statuses, err := statusFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to find status by board, got nil")
		}
		if len(statuses) != 0 {
			t.Errorf("Expected empty slice, got %v", statuses)
		}
	})

	t.Run("should return error on invalid id", func(t *testing.T) {
		statusFinder := NewStatusFinderByBoard(*slog.Default(), nil, nil)
		statuses, err := statusFinder.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Expected error on invalid id, got nil")
		}
		if len(statuses) != 0 {
			t.Errorf("Expected empty slice, got %v", statuses)
		}
	})
}