package local

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
)

type TaskRepository struct {
	logger slog.Logger
	data map[valueobject.Id]model.Task
}

func NewTaskRepository(logger slog.Logger) TaskRepository {
	return TaskRepository{logger, make(map[valueobject.Id]model.Task)}
}

func (tr TaskRepository) Find(ctx context.Context, taskId valueobject.Id) (pkg.Optional[model.Task], error) {
	tr.logger.Info("TaskRepository - Find - Params: ", slog.Any("taskId", taskId))
	task, ok := tr.data[taskId]
	if !ok {
		return pkg.EmptyOptional[model.Task](), nil
	}

	return pkg.NewOptional(task), nil
}

func (tr TaskRepository) FindByTeam(ctx context.Context, boardId valueobject.Id) (pkg.Optional[[]model.Task], error) {
	tr.logger.Info("TaskRepository - FindByTeam - Params: ", slog.Any("boardId", boardId))
	var tasks []model.Task
	for _, task := range tr.data {
		if task.BoardId == boardId {
			tasks = append(tasks, task)
		}
	}
	
	if len(tasks) == 0 {
		return pkg.EmptyOptional[[]model.Task](), nil
	}

	return pkg.NewOptional(tasks), nil
}

func (tr TaskRepository) Save(ctx context.Context, task model.Task) error {
	tr.logger.Info("TaskRepository - Save - Params: ", slog.Any("task", task))
	tr.data[task.Id] = task
	return nil
}

func (tr TaskRepository) Delete(ctx context.Context, taskId valueobject.Id) error {
	tr.logger.Info("TaskRepository - Delete - Params: ", slog.Any("taskId", taskId))
	delete(tr.data, taskId)
	return nil
}

func (tr TaskRepository) Update(ctx context.Context, task model.Task) error {
	tr.logger.Info("TaskRepository - Update - Params: ", slog.Any("task", task))
	tr.data[task.Id] = task
	return nil
}