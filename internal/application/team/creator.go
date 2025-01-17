package team

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/user"
)

type TeamCreator struct {
	logger slog.Logger
	repository repository.TeamRepository
	userFinder user.UserFinder
}

func NewTeamCreator(
		logger slog.Logger, 
		teamRepository repository.TeamRepository,
		userRepository repository.UserRepository,
	) TeamCreator {
	return TeamCreator{
		logger: logger, 
		repository: teamRepository, 
		userFinder:  user.NewUserFinder(logger, userRepository),
	}
}

func (tc TeamCreator) Run(ctx context.Context, title, owner, description string) error {
	tc.logger.Info(
		"TeamCreator - Run - Params: ", 
		slog.Any("title", title), slog.Any("owner", owner), 
		slog.Any("description", 
		description),
	)

	team, err := model.NewTeam(title, description, owner)
	if err != nil {
		tc.logger.Info("TeamCreator - Run - Error has ocurred trying to instance new team")
		return err
	}

	_, errUser := tc.userFinder.Run(ctx, owner)
	if errUser != nil {
		tc.logger.Info("TeamCreator - Run - Error has ocurred trying to find user", slog.Any("err", errUser))
		return errUser
	}

	return tc.repository.Save(ctx, team)
}