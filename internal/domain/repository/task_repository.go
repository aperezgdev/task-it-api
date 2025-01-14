package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
)

type TaskRepository interface {
	save(ctx context.Context, task model.Task) error
	delete(ctx context.Context, task model.Task) error
	update(ctx context.Context, task model.Task) error
}