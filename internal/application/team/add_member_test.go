package team

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

func TestAddMember(t *testing.T) {
	t.Parallel()

	t.Run("should add a member to a team without error", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		userRepository := &repository.MockUserRepository{}
		teamAddMember := NewTeamAddMember(*slog.Default(), teamRepository, userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		teamRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := teamAddMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
		teamRepository.AssertNumberOfCalls(t, "Update", 1)
	})

	t.Run("should return error on not existing team", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		userRepository := &repository.MockUserRepository{}
		teamAddMember := NewTeamAddMember(*slog.Default(), teamRepository, userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		teamRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := teamAddMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on not existing team")
		}		

		if !errors.Is(err, domain_errors.ErrNotExist) {
			t.Errorf("Error should be ErrNotExist")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 1)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
		teamRepository.AssertNumberOfCalls(t, "Update", 0)
	})

	t.Run("should return error on not existing user", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		userRepository := &repository.MockUserRepository{}
		teamAddMember := NewTeamAddMember(*slog.Default(), teamRepository, userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), nil)
		teamRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := teamAddMember.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9", "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on not existing user")
		}		

		if !errors.Is(err, domain_errors.ErrNotExist) {
			t.Errorf("Error should be ErrNotExist")
		}

		teamRepository.AssertNumberOfCalls(t, "Find", 0)
		userRepository.AssertNumberOfCalls(t, "Find", 1)
		teamRepository.AssertNumberOfCalls(t, "Update", 0)
	})
}