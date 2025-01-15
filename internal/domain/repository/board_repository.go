package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type BoardRepository interface {
	Save(ctx context.Context, board model.Board) error
	Delete(ctx context.Context, boardId valueobject.Id) error
	Update(ctx context.Context, board model.Board) error
}