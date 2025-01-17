package team

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestTeamCreator(t *testing.T) {
	t.Parallel()

	t.Run("should create a valid team without error", func(t *testing.T) {
		mockTeamRepository := repository.MockTeamRepository{}
		mockUserRepository := repository.MockUserRepository{}
		teamCreator := NewTeamCreator(*slog.Default(), &mockTeamRepository, &mockUserRepository)
		mockUserRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil).Once()
		mockTeamRepository.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
		err := teamCreator.Run(context.Background(), "title", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "description")
	
		if err != nil {
			t.Errorf("Error shouldnt happened on a valid team")
		}

		mockTeamRepository.AssertNumberOfCalls(t, "Save", 1)
		mockUserRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on a not exising user", func(t *testing.T) {
		mockTeamRepository := repository.MockTeamRepository{}
		mockUserRepository := repository.MockUserRepository{}
		teamCreator := NewTeamCreator(*slog.Default(), &mockTeamRepository, &mockUserRepository)
		mockUserRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), nil).Once()
		err := teamCreator.Run(context.Background(), "title", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "description")
	
		if err == nil {
			t.Errorf("Error should happened on a not existing user")
		}

		mockTeamRepository.AssertNumberOfCalls(t, "Save", 0)
		mockUserRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on an invalid team", func(t *testing.T) {
		mockTeamRepository := repository.MockTeamRepository{}
		mockUserRepository := repository.MockUserRepository{}
		teamCreator := NewTeamCreator(*slog.Default(), &mockTeamRepository, &mockUserRepository)
		err := teamCreator.Run(context.Background(), "", "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "description")
	
		if err == nil {
			t.Errorf("Error should happened on an invalid team")
		}

		mockTeamRepository.AssertNumberOfCalls(t, "Save", 0)
		mockUserRepository.AssertNumberOfCalls(t, "Find", 0)
	})
}