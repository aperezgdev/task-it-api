package board

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

func TestBoardRemover(t *testing.T) {
	t.Run("should remove board", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)

		boardRemover := NewBoardRemover(*slog.Default(), boardRepository)
		err := boardRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("should return error trying to find board", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), nil)

		boardRemover := NewBoardRemover(*slog.Default(), boardRepository)
		err := boardRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to find board, got nil")
		}
	})

	t.Run("should return error trying to delete board", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.ErrNotExist)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)

		boardRemover := NewBoardRemover(*slog.Default(), boardRepository)
		err := boardRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to delete board, got nil")
		}
	})

	t.Run("should return error on invalid id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		boardFinder := &repository.MockBoardRepository{}
		boardFinder.On("Run", mock.Anything, mock.Anything).Return(model.Board{}, nil)

		boardRemover := NewBoardRemover(*slog.Default(), boardRepository)
		err := boardRemover.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Expected error on invalid id, got nil")
		}
	})
}