package user

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

func TestUserFinder(t *testing.T) {
	t.Parallel()

	mockRepository := &repository.MockUserRepository{}
	userFinder := NewUserFinder(*slog.Default(), mockRepository)

	t.Run("should find an existing user", func(t *testing.T) {
		mockRepository.On("Find", mock.Anything, mock.Anything).Once().Return(pkg.NewOptional(model.User{}), nil)
		_, err := userFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on a existing user")
		}

		mockRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return not exist error on not an existing user", func(t *testing.T) {
		mockRepository.On("Find", mock.Anything, mock.Anything).Once().Return(pkg.EmptyOptional[model.User](), nil)
		_, err := userFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")

		if err == nil {
			t.Errorf("Error should happened on a not existing user")
		}

		if !errors.Is(err, domain_errors.ErrNotExist) {
			t.Errorf("Error should be ErrNotExist")
		}

		mockRepository.AssertNumberOfCalls(t, "Find", 1)
	})
}