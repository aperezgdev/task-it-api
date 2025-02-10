package task

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/team"
)

type TaskFinderByTeam struct {
	logger slog.Logger
	repository repository.TaskRepository
	teamFinder team.TeamFinder
}

func NewTaskFinderByTeam(logger slog.Logger, repository repository.TaskRepository, teamRepository repository.TeamRepository) TaskFinderByTeam {
	return TaskFinderByTeam{logger, repository, team.NewTeamFinder(logger, teamRepository)}
}

func (tf TaskFinderByTeam) Run(ctx context.Context, teamId string) ([]model.Task, error) {
	tf.logger.Info("TaskFinderByTeam - Run - Params: ", slog.Any("teamId", teamId))
	team, errTeam := tf.teamFinder.Run(ctx, teamId)
	if errTeam != nil {
		tf.logger.Info("TaskFinderByTeam - Run - Team not found")
		return []model.Task{}, errTeam
	}

	taskOptional, err := tf.repository.FindByTeam(ctx, team.Id)
	if err != nil {
		return []model.Task{}, err
	}

	if !taskOptional.IsPresent {
		return []model.Task{}, nil
	}

	return taskOptional.Value, nil
}