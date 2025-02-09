package task

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

func TestTaskRemover(t *testing.T) {
	t.Run("should remove task", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)

		taskRemover := NewTaskRemover(*slog.Default(), taskRepository)
		err := taskRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("should return error trying to find task", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Task](), errors.ErrNotExist)

		taskRemover := NewTaskRemover(*slog.Default(), taskRepository)
		err := taskRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to find task, got nil")
		}
	})

	t.Run("should return error trying to delete task", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.ErrNotExist)
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)

		taskRemover := NewTaskRemover(*slog.Default(), taskRepository)
		err := taskRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to delete task, got nil")
		}
	})

	t.Run("should return error on invalid id", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		taskFinder := &repository.MockTaskRepository{}
		taskFinder.On("Run", mock.Anything, mock.Anything).Return(model.Task{}, nil)

		taskRemover := NewTaskRemover(*slog.Default(), taskRepository)
		err := taskRemover.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Expected error on invalid id, got nil")
		}
	})
}