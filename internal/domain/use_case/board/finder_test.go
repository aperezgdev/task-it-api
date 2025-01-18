package board

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	domain_errors "github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestFinder(t *testing.T) {
	t.Parallel()

	t.Run("should find existing board", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardFinder := NewBoardFinder(*slog.Default(), boardRepository)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		_, err := boardFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		
		if err != nil {
			t.Errorf("Error shouldnt happened on existing board")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return an error on not existing board", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardFinder := NewBoardFinder(*slog.Default(), boardRepository)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), nil)
		_, err := boardFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		
		if err == nil {
			t.Errorf("Error should happened on not existing board")
		}

		if !errors.Is(err, domain_errors.ErrNotExist) {
			t.Errorf("Error should be not exist error")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return an error on invalid id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardFinder := NewBoardFinder(*slog.Default(), boardRepository)
		_, err := boardFinder.Run(context.Background(), "1")
		
		if err == nil {
			t.Errorf("Error shouldnt on invalid in")
		}

		_, ok := err.(domain_errors.ValidationError)
		if !ok {
			t.Errorf("Error should be not exist error")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 0)
	})
}