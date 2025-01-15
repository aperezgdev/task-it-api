package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type StatusRepository interface {
	Save(ctx context.Context, status model.Status) error
	Delete(ctx context.Context, statusId valueobject.Id) error
	Update(ctx context.Context, status model.Status) error
}