package task

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/task"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type TaskRemover struct {
	logger slog.Logger
	repository repository.TaskRepository
	finder task.TaskFinder
}

func NewTaskRemover(logger slog.Logger, repository repository.TaskRepository) TaskRemover {
	return TaskRemover{
		logger: logger,
		repository: repository,
		finder: task.NewTaskFinder(logger, repository),
	}
}

func (tr *TaskRemover) Run(ctx context.Context, id string) error {
	taskId, errId := valueobject.ValidateId(id)
	if errId != nil {
		tr.logger.Info("TaskRemover - Run - Id not valid")
		return errId
	}
	
	_, err := tr.finder.Run(ctx, id)
	if err != nil {
		tr.logger.Info("TaskRemover - Run - Error trying to find task")
		return err
	}

	err = tr.repository.Delete(ctx, taskId)
	if err != nil {
		tr.logger.Info("TaskRemover - Run - Error trying to delete task")
		return err
	}

	return nil
}