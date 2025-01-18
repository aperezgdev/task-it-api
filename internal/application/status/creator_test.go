package status

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestCreator(t *testing.T) {
	t.Parallel()

	t.Run("should create a status on valid params", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		creator := NewStatusCreator(*slog.Default(), statusRepository, boardRepository)

		err := creator.Run(context.Background(), "titulo", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", []string{"01946ba3-ee73-76e6-83a9-33f87a35d6e9"}, []string{})
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("should return error on invalid board id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), nil)
		statusRepository := &repository.MockStatusRepository{}
		creator := NewStatusCreator(*slog.Default(), statusRepository, boardRepository)

		err := creator.Run(context.Background(), "titulo", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", []string{"01946ba3-ee73-76e6-83a9-33f87a35d6e9"}, []string{})
		if err == nil {
			t.Errorf("Error should happened on valid params")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Save", 0)
	})

	t.Run("should return error on invalid params", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		statusRepository := &repository.MockStatusRepository{}
		creator := NewStatusCreator(*slog.Default(), statusRepository, boardRepository)

		err := creator.Run(context.Background(), "", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", []string{"01946ba3-ee73-76e6-83a9-33f87a35d6e9"}, []string{})
		if err == nil {
			t.Errorf("Error should happened on invalid params")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Save", 0)
	})
}