package task

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

func TestUpdater(t *testing.T) {
	t.Parallel()

	t.Run("should update a task without error", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskUpdater := NewTaskUpdater(*slog.Default(), taskRepository)
		taskRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := taskUpdater.Run(context.Background(), model.Task{})
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		taskRepository.AssertNumberOfCalls(t, "Update", 1)
	})
}