package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
)

type TaskRepository interface {
	Save(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, task model.Task) error
	Update(ctx context.Context, task model.Task) error
}