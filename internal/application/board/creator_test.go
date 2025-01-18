package board

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

	t.Run("should create a valid board without error", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		userRepository := &repository.MockUserRepository{}
		teamRepository := &repository.MockTeamRepository{}
		boardRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		creator := NewBoardCreator(*slog.Default(), boardRepository, userRepository, teamRepository)
		err := creator.Run(context.Background(), "title", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		boardRepository.AssertNumberOfCalls(t, "Save", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on invalid team id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		userRepository := &repository.MockUserRepository{}
		teamRepository := &repository.MockTeamRepository{}
		boardRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), nil)
		creator := NewBoardCreator(*slog.Default(), boardRepository, userRepository, teamRepository)
		err := creator.Run(context.Background(), "title", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid team id")
		}

		boardRepository.AssertNumberOfCalls(t, "Save", 0)
		userRepository.AssertNumberOfCalls(t, "Find", 0)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on invalid user id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		userRepository := &repository.MockUserRepository{}
		teamRepository := &repository.MockTeamRepository{}
		boardRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), nil)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		creator := NewBoardCreator(*slog.Default(), boardRepository, userRepository, teamRepository)
		err := creator.Run(context.Background(), "title", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid user id")
		}

		boardRepository.AssertNumberOfCalls(t, "Save", 0)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on invalid board params", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		userRepository := &repository.MockUserRepository{}
		teamRepository := &repository.MockTeamRepository{}
		creator := NewBoardCreator(*slog.Default(), boardRepository, userRepository, teamRepository)
		err := creator.Run(context.Background(), "", "description", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on invalid board params")
		}

		boardRepository.AssertNumberOfCalls(t, "Save", 0)	
		userRepository.AssertNumberOfCalls(t, "Find", 0)
		teamRepository.AssertNumberOfCalls(t, "Find", 0)
	})

}