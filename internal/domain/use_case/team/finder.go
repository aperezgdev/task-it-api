package team

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type TeamFinder struct {
	logger slog.Logger
	repository repository.TeamRepository
}

func NewTeamFinder(logger slog.Logger, repository repository.TeamRepository) TeamFinder {
	return TeamFinder{logger, repository}	
}

func (tf *TeamFinder) Run(ctx context.Context,idTeam string) (model.Team, error) {
	tf.logger.Info("TeamFinder - Run - Params: ", slog.Any("idTeam", idTeam))
	id, errId := valueobject.ValidateId(idTeam)
	if errId != nil {
		tf.logger.Info("TeamFinder - Run - Team id was invalid")
		return model.Team{}, errId
	}

	teamOptional, err := tf.repository.Find(ctx, id)
	if err != nil {
		tf.logger.Error("TeamFinder - Run - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return model.Team{}, err
	}

	if !teamOptional.IsPresent {
		tf.logger.Info("TeamFinder - Run - Team not exist")
		return model.Team{}, errors.ErrNotExist
	}

	return teamOptional.Value, nil
}