package user

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type UserFinder struct {
	logger slog.Logger
	repository repository.UserRepository
}

func NewUserFinder(logger slog.Logger, repository repository.UserRepository) UserFinder {
	return UserFinder{logger, repository}
}

func (uf *UserFinder) Run(ctx context.Context, userId string) (model.User, error) {
	uf.logger.Info("UserFinder - Run - Params: ", slog.Any("userId", userId))
	userIdVO, errId := valueobject.ValidateId(userId)
	if errId != nil {
		uf.logger.Info("UserFinder - Run - User Id is not valid")
		return model.User{}, errId
	}

	userOptional, err := uf.repository.Find(ctx, userIdVO)
	if err != nil {
		uf.logger.Error("UserFinder - Run - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return model.User{}, err
	}

	if !userOptional.IsPresent {
		uf.logger.Info("UserFinder - Run - User not exist")
		return model.User{}, errors.ErrNotExists
	}

	return userOptional.Value, nil
}