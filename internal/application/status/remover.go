package status

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/status"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type StatusRemover struct {
	logger slog.Logger
	repository repository.StatusRepository
	finder status.StatusFinder
}

func NewStatusRemover(logger slog.Logger, repository repository.StatusRepository) StatusRemover {
	return StatusRemover{
		logger: logger,
		repository: repository,
		finder: status.NewStatusFinder(logger, repository),
	}
}

func (sr *StatusRemover) Run(ctx context.Context, id string) error {
	statusId, errId := valueobject.ValidateId(id)
	if errId != nil {
		sr.logger.Info("StatusRemover - Run - Id not valid")
		return errId
	}
	
	_, err := sr.finder.Run(ctx, id)
	if err != nil {
		sr.logger.Info("StatusRemover - Run - Error trying to find status")
		return err
	}

	err = sr.repository.Delete(ctx, statusId)
	if err != nil {
		sr.logger.Info("StatusRemover - Run - Error trying to delete status")
		return err
	}

	return nil
}