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

func TestTaskFinderByTeam(t *testing.T) {
	t.Parallel()

	t.Run("should return task by team", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.NewOptional([]model.Task{}), nil)
		teamRepository := &repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		taskFinder := NewTaskFinderByTeam(*slog.Default(), taskRepository, teamRepository)

		tasks, err := taskFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(tasks) != 0 {
			t.Errorf("Expected empty array, got %v", tasks)
		}

		taskRepository.AssertNumberOfCalls(t, "FindByTeam", 1)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on invalid team id", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[[]model.Task](), nil)
		teamRepository := &repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), nil)
		taskFinder := NewTaskFinderByTeam(*slog.Default(), taskRepository, teamRepository)

		tasks, err := taskFinder.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Expected error on invalid team id, got nil")
		}

		if len(tasks) != 0 {
			t.Errorf("Expected empty array, got %v", tasks)
		}

		taskRepository.AssertNumberOfCalls(t, "FindByTeam", 0)
		teamRepository.AssertNumberOfCalls(t, "Find", 0)
	})

	t.Run("should return error on not existing team", func(t *testing.T) {
		taskRepository := &repository.MockTaskRepository{}
		taskRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[[]model.Task](), nil)
		teamRepository := &repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), nil)
		taskFinder := NewTaskFinderByTeam(*slog.Default(), taskRepository, teamRepository)

		tasks, err := taskFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error on task not found, got nil")
		}

		if len(tasks) != 0 {
			t.Errorf("Expected empty array, got %v", tasks)
		}

		taskRepository.AssertNumberOfCalls(t, "FindByTeam", 0)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})
}