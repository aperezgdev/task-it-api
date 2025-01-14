package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type BoardRepository interface {
	save(ctx context.Context, board model.Board) error
	delete(ctx context.Context, boardId valueobject.Id) error
	update(ctx context.Context, board model.Board) error
}