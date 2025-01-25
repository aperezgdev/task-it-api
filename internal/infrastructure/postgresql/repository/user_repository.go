package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/postgresql/sqlc"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *sqlc.Queries
	logger  slog.Logger
}

func NewUserRepository(logger slog.Logger, queries *sqlc.Queries) UserRepository {
	return UserRepository{queries, logger}
}

func (ur UserRepository) Find(ctx context.Context, userId valueobject.Id) (pkg.Optional[model.User], error) {
	user, err := ur.queries.FindUser(ctx, uuid.UUID(userId))
	if err != nil {
		ur.logger.Error("UserRepository - Find - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return pkg.EmptyOptional[model.User](), err
	}

	return pkg.NewOptional(model.User{
		Id: valueobject.Id(user.ID),
		Email: valueobject.Email(user.Email),
		CreatedAt: valueobject.CreatedAt(user.CreatedAt.Time),
	}), nil
}

func (ur UserRepository) Save(ctx context.Context, user model.User) error {
	_, err := ur.queries.SaveUser(ctx, sqlc.SaveUserParams{
		ID:        uuid.UUID(user.Id),
		Email:     string(user.Email),
		CreatedAt: pgtype.Timestamp{Time: time.Time(user.CreatedAt)},
	})
	if err != nil {
		ur.logger.Error("UserRepository - Save - Error has ocurred trying to save data to repository", slog.Any("error", err))
		return err
	}

	return nil
}


