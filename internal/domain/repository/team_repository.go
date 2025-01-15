package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
)

type TeamRepository interface {
	Save(ctx context.Context, team model.Team) error
	Delete(ctx context.Context, team model.Team) error
	Update(ctx context.Context, team model.Team) error
}
