package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
)

type TeamRepository interface {
	save(ctx context.Context, team model.Team) error
	delete(ctx context.Context, team model.Team) error
	update(ctx context.Context, team model.Team) error
}
