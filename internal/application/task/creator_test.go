package task

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestTaskCreator(t *testing.T) {
	t.Parallel()

	t.Run("should create a task on valid params", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository := &repository.MockUserRepository{}
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		creator := NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)

		err := creator.Run(context.Background(), "title", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
		taskRepository.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("should return error on invalid board id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), nil)
		userRepository := &repository.MockUserRepository{}
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		taskRepository := &repository.MockTaskRepository{}
		creator := NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)

		err := creator.Run(context.Background(), "title", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid board id")
		}

		boardRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 0)
		taskRepository.AssertNumberOfCalls(t, "Save", 0)
	})

	t.Run("should return error on invalid user id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository := &repository.MockUserRepository{}
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), nil)
		taskRepository := &repository.MockTaskRepository{}
		creator := NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)

		err := creator.Run(context.Background(), "title", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid user id")
		}
	})

	t.Run("should return error on invalid task params", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository := &repository.MockUserRepository{}
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		taskRepository := &repository.MockTaskRepository{}
		creator := NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)

		err := creator.Run(context.Background(), "", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid task params")
		}
	})
}