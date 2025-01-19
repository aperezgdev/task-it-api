package task

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type TaskFinder struct {
	logger slog.Logger
	repository repository.TaskRepository
}

func NewTaskFinder(logger slog.Logger, repository repository.TaskRepository) TaskFinder {
	return TaskFinder{logger, repository}
}

func (tf *TaskFinder) Run(ctx context.Context, taskId string) (model.Task, error) {
	tf.logger.Info("TaskFinder - Run - Params: ", slog.Any("taskId", taskId))
	id, errId := valueobject.ValidateId(taskId)
	if errId != nil {
		tf.logger.Info("TaskFinder - Run - Task id was invalid")
		return model.Task{}, errId
	}

	taskOptional, err := tf.repository.Find(ctx, id)
	if err != nil {
		tf.logger.Error("TaskFinder - Run - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return model.Task{}, err
	}

	if !taskOptional.IsPresent {
		tf.logger.Info("TaskFinder - Run - Task not exist")
		return model.Task{}, errors.ErrNotExist
	}

	return taskOptional.Value, nil
}