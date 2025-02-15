package local

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
)

type UserRepository struct {
	logger slog.Logger
	data map[valueobject.Id]model.User
}

func NewUserRepository(logger slog.Logger) UserRepository {
	return UserRepository{logger, make(map[valueobject.Id]model.User)}
}

func (ur UserRepository) Find(ctx context.Context, userId valueobject.Id) (pkg.Optional[model.User], error) {
	ur.logger.Info("UserRepository - Find - Params: ", slog.Any("userId", userId))
	user, ok := ur.data[userId]
	if !ok {
		return pkg.EmptyOptional[model.User](), nil
	}

	return pkg.NewOptional(user), nil
}

func (ur UserRepository) Save(ctx context.Context, user model.User) error {
	ur.logger.Info("UserRepository - Save - Params: ", slog.Any("user", user))
	ur.data[user.Id] = user
	return nil
}

func (ur UserRepository) Delete(ctx context.Context, userId valueobject.Id) error {
	ur.logger.Info("UserRepository - Delete - Params: ", slog.Any("userId", userId))
	delete(ur.data, userId)
	return nil
}

func (ur UserRepository) Update(ctx context.Context, user model.User) error {
	ur.logger.Info("UserRepository - Update - Params: ", slog.Any("user", user))
	ur.data[user.Id] = user
	return nil
}