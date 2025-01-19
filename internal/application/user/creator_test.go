package user

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

func TestCreator(t *testing.T) {
	t.Parallel()

	t.Run("should create a valid user without error", func(t *testing.T) {
		userRepository := &repository.MockUserRepository{}
		userRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		creator := NewUserCreator(*slog.Default(), userRepository)
		err := creator.Run(context.Background(), "aperezgdev@example.com")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		userRepository.AssertNumberOfCalls(t, "Save", 1)
	})

	t.Run("should return error on invalid email", func(t *testing.T) {
		userRepository := &repository.MockUserRepository{}
		creator := NewUserCreator(*slog.Default(), userRepository)
		err := creator.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Error should happened on invalid email")
		}

		userRepository.AssertNumberOfCalls(t, "Save", 0)
	})
}