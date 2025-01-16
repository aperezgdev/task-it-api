package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
)

type UserRepository interface {
	Find(ctx context.Context, userId valueobject.Id) (pkg.Optional[model.User], error)
	Save(ctx context.Context, user model.User) error
}