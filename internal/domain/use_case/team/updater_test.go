package team

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

	t.Run("should update a team without error", func(t *testing.T) {
		teamRepository := &repository.MockTeamRepository{}
		teamUpdater := NewTeamUpdater(*slog.Default(), teamRepository)
		teamRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := teamUpdater.Run(context.Background(), model.Team{})
		if err != nil {
			t.Errorf("Error shouldnt happened on valid params")
		}

		teamRepository.AssertNumberOfCalls(t, "Update", 1)
	})
}