package status

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestStatusRemover(t *testing.T) {
	t.Run("should remove status", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		statusRemover := NewStatusRemover(*slog.Default(), statusRepository)
		err := statusRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("should return error trying to find status", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Status](), errors.ErrNotExist)

		statusRemover := NewStatusRemover(*slog.Default(), statusRepository)
		err := statusRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to find status, got nil")
		}
	})

	t.Run("should return error trying to delete status", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.ErrNotExist)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		statusRemover := NewStatusRemover(*slog.Default(), statusRepository)
		err := statusRemover.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error trying to delete status, got nil")
		}
	})

	t.Run("should return error on invalid id", func(t *testing.T) {
		statusRepository := &repository.MockStatusRepository{}
		statusRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		statusFinder := &repository.MockStatusRepository{}
		statusFinder.On("Run", mock.Anything, mock.Anything).Return(model.Status{}, nil)

		statusRemover := NewStatusRemover(*slog.Default(), statusRepository)
		err := statusRemover.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Expected error on invalid id, got nil")
		}
	})
}