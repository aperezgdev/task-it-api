package team

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
)

type TeamUpdater struct {
	logger slog.Logger
	repository repository.TeamRepository
	teamFinder TeamFinder
}

func NewTeamUpdater(logger slog.Logger, repository repository.TeamRepository) TeamUpdater {
	return TeamUpdater{logger, repository, NewTeamFinder(logger, repository)}
}

func (tu *TeamUpdater) Run(ctx context.Context, team model.Team) error {
	tu.logger.Info("TeamUpdater - Run - Params: ", slog.Any("team", team))

	return tu.repository.Update(ctx, team)
}