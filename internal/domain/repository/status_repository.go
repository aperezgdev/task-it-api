package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type StatusRepository interface {
	save(ctx context.Context, status model.Status) error
	delete(ctx context.Context, statusId valueobject.Id) error
	update(ctx context.Context, status model.Status) error
}