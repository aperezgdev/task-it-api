package team

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestRemoverMember(t *testing.T) {
	t.Parallel()

	t.Run("should remove a member from a team", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		userRepository := &repository.MockUserRepository{}
		removerMember := NewRemoverMember(*slog.Default(), teamRepository, userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		err := removerMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on a valid team")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on not existing team", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		userRepository := &repository.MockUserRepository{}
		removerMember := NewRemoverMember(*slog.Default(), teamRepository, userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		err := removerMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on a not existing team")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 0)
	})

	t.Run("should return error on not existing user", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		userRepository := &repository.MockUserRepository{}
		removerMember := NewRemoverMember(*slog.Default(), teamRepository, userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), nil)
		err := removerMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on a not existing user")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on user already removed", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}			
		userRepository := &repository.MockUserRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{Members: []valueobject.Id{valueobject.Id([16]byte{1})}}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		removerMember := NewRemoverMember(*slog.Default(), teamRepository, userRepository)
		err := removerMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on a user already removed")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
	})
}