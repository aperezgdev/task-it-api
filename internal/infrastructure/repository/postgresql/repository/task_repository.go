package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/repository/postgresql/sqlc"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskRepository struct {
	queries *sqlc.Queries
	logger  slog.Logger
}

func NewTaskRepository(logger slog.Logger, queries *sqlc.Queries) TaskRepository {
	return TaskRepository{queries, logger}
}

func (tr TaskRepository) Find(ctx context.Context, taskId valueobject.Id) (pkg.Optional[model.Task], error) {
	task, err := tr.queries.FindTask(ctx, uuid.UUID(taskId))
	if err != nil {
		tr.logger.Error("TaskRepository - Find - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return pkg.EmptyOptional[model.Task](), err
	}

	return pkg.NewOptional(model.Task{
		Id: valueobject.Id(task.ID),
		Title: valueobject.Title(task.Title),
		Description: valueobject.Description(task.Description),
		Creator: valueobject.Id(task.Creator),
		Asigned: valueobject.Id(task.Asigned),
		StatusId: valueobject.Id(task.StatusID),
		CreatedAt: valueobject.CreatedAt(task.CreatedAt.Time),
	}), nil
}

func (tr TaskRepository) Save(ctx context.Context, task model.Task) error {
	_, err := tr.queries.SaveTask(ctx, sqlc.SaveTaskParams{
		ID:          uuid.UUID(task.Id),
		Title:       string(task.Title),
		Description: string(task.Description),
		Creator:     uuid.UUID(task.Creator),
		Asigned:     uuid.UUID(task.Asigned),
		StatusID:    uuid.UUID(task.StatusId),
		CreatedAt:   pgtype.Timestamp{Time: time.Time(task.CreatedAt)},
	})
	if err != nil {
		tr.logger.Error("TaskRepository - Save - Error has ocurred trying to save data to repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (tr TaskRepository) Delete(ctx context.Context, taskId valueobject.Id) error {
	err := tr.queries.DeleteTask(ctx, uuid.UUID(taskId))
	if err != nil {
		tr.logger.Error("TaskRepository - Delete - Error has ocurred trying to delete data from repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (tr TaskRepository) Update(ctx context.Context, task model.Task) error {
	err := tr.queries.UpdateTask(ctx, sqlc.UpdateTaskParams{
		ID:          uuid.UUID(task.Id),
		Title:       string(task.Title),
		Description: string(task.Description),
		Creator:     uuid.UUID(task.Creator),
		Asigned:     uuid.UUID(task.Asigned),
		StatusID:    uuid.UUID(task.StatusId),
		CreatedAt:  pgtype.Timestamp{Time: time.Time(task.CreatedAt)},
	})
	if err != nil {
		tr.logger.Error("TaskRepository - Update - Error has ocurred trying to update data from repository", slog.Any("error", err))
		return err
	}

	return nil
}