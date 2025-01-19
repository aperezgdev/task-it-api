package task

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
)

type TaskUpdater struct {
	logger slog.Logger
	repository repository.TaskRepository
}

func NewTaskUpdater(logger slog.Logger, repository repository.TaskRepository) TaskUpdater {
	return TaskUpdater{logger, repository}
}

func (tu *TaskUpdater) Run(ctx context.Context, task model.Task) error {
	tu.logger.Info("TaskUpdater - Run - Params: ", slog.Any("task", task))

	return tu.repository.Update(ctx, task)
}