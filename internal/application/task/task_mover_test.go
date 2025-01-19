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

func TestMover(t *testing.T) {
	t.Parallel()

	t.Run("should move a task to a status without error", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)
		taskRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		creator := NewTaskMover(*slog.Default(), taskRepository, statusRepository)

		err := creator.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		taskRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Find", 1)
		taskRepository.AssertNumberOfCalls(t, "Update", 1)
	})

	t.Run("should return error on invalid task id", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Task](), nil)		
		statusRepository := &repository.MockStatusRepository{}
		creator := NewTaskMover(*slog.Default(), taskRepository, statusRepository)

		err := creator.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid task id")
		}

		taskRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Find", 0)
	})

	t.Run("should return error on invalid status id", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Status](), nil)
		taskRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		creator := NewTaskMover(*slog.Default(), taskRepository, statusRepository)

		err := creator.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid status id")
		}

		taskRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Find", 1)
		taskRepository.AssertNumberOfCalls(t, "Update", 0)
	})
}