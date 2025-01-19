package status

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type StatusFinder struct {
	logger slog.Logger
	repository repository.StatusRepository
}

func NewStatusFinder(logger slog.Logger, repository repository.StatusRepository) StatusFinder {
	return StatusFinder{logger, repository}
}

func (sf *StatusFinder) Run(ctx context.Context, statusId string) (model.Status, error) {
	sf.logger.Info("StatusFinder - Run - Params: ", slog.Any("statusId", statusId))
	id, errId := valueobject.ValidateId(statusId)
	if errId != nil {
		sf.logger.Info("StatusFinder - Run - Status id was invalid")
		return model.Status{}, errId
	}

	statusOptional, err := sf.repository.Find(ctx, id)
	if err != nil {
		sf.logger.Error("StatusFinder - Run - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return model.Status{}, err
	}

	if !statusOptional.IsPresent {
		sf.logger.Info("StatusFinder - Run - Status not exist")
		return model.Status{}, errors.ErrNotExist
	}

	return statusOptional.Value, nil
}