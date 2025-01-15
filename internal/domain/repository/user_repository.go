package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) error
}