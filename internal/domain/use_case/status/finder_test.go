package status

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

func TestFinder(t *testing.T) {
	t.Parallel()

	t.Run("should find existing status", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusFinder := NewStatusFinder(*slog.Default(), statusRepository)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)
		_, err := statusFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on a existing status")
		}

		statusRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return not exist error on not an existing status", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusFinder := NewStatusFinder(*slog.Default(), statusRepository)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Status](), nil)		
		_, err := statusFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Error should happened on a not existing status")
		}

		if !errors.Is(err, domain_errors.ErrNotExist) {
			t.Errorf("Error should be not exist error")
		}		

		statusRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return an error on invalid id", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusFinder := NewStatusFinder(*slog.Default(), statusRepository)
		_, err := statusFinder.Run(context.Background(), "1")
		if err == nil {
			t.Errorf("Error shouldnt on invalid in")
		}

		_, ok := err.(domain_errors.ValidationError)
		if !ok {
			t.Errorf("Error should be not exist error")
		}

		statusRepository.AssertNumberOfCalls(t, "Find", 0)
	})
}