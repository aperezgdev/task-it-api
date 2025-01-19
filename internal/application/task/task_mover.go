package task

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/status"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/task"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type TaskMover struct {
	logger slog.Logger
	taskUpdater  task.TaskUpdater
	taskFinder task.TaskFinder
	statusFinder status.StatusFinder
}

func NewTaskMover(logger slog.Logger, taskRepository repository.TaskRepository, statusRepository repository.StatusRepository) TaskMover {
	return TaskMover{
		logger: logger, 
		taskUpdater:  task.NewTaskUpdater(logger, taskRepository),
		taskFinder:  task.NewTaskFinder(logger, taskRepository),
		statusFinder: status.NewStatusFinder(logger, statusRepository),
	}
}

func (tm TaskMover) Run(ctx context.Context, taskId, statusId string) error {
	tm.logger.Info("TaskMover - Run - Params: ", slog.Any("taskId", taskId), slog.Any("statusId", statusId))
	task, errTask := tm.taskFinder.Run(ctx, taskId)
	if errTask != nil {
		tm.logger.Info("TaskMover - Run - Task not found")
		return errTask
	}

	_, errStatus := tm.statusFinder.Run(ctx, statusId)
	if errStatus != nil {
		tm.logger.Info("TaskMover - Run - Status not found")
		return errStatus
	}

	taskIdBytes := []byte(taskId)
	var taskIdArray [16]byte
	copy(taskIdArray[:], taskIdBytes)

	task.StatusId = valueobject.Id(taskIdArray)

	return tm.taskUpdater.Run(ctx, task)
}