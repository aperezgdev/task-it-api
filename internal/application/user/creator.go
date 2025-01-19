package user

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
)

type UserCreator struct {
	logger slog.Logger
	repository repository.UserRepository
}

func NewUserCreator(logger slog.Logger, repository repository.UserRepository) UserCreator {
	return UserCreator{logger, repository}
}

func (uc *UserCreator) Run(ctx context.Context, email string) error {
	uc.logger.Info("UserCreator - Run - Params: ", slog.Any("email", email))
	user, err := model.NewUser(email)
	if err != nil {
		uc.logger.Info("UserCreator - Run - Error trying to create user")
		return err
	}

	return uc.repository.Save(ctx, user)
}