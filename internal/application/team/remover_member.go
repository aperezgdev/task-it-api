package team

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/team"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/user"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type RemoverMember struct {
	logger slog.Logger
	teamUpdater team.TeamUpdater
	userFinder user.UserFinder
	teamFinder team.TeamFinder
}

func NewRemoverMember(logger slog.Logger, teamRepository repository.TeamRepository, userRepository repository.UserRepository) RemoverMember {
	return RemoverMember{logger, team.NewTeamUpdater(logger, teamRepository), user.NewUserFinder(logger, userRepository), team.NewTeamFinder(logger, teamRepository)}
}

func (rm RemoverMember) Run(ctx context.Context, teamId, userId string) error {
	rm.logger.Info("RemoverMember - Run - Params: ", slog.Any("teamId", teamId), slog.Any("userId", userId))
	team, errTeam := rm.teamFinder.Run(ctx, teamId)
	if errTeam != nil {
		rm.logger.Info("RemoverMember - Run - Team not found")
		return errTeam
	}

	_, errUser := rm.userFinder.Run(ctx, userId)
	if errUser != nil {
		rm.logger.Info("RemoverMember - Run - User not found")
		return errUser
	}

	size := len(team.Members)
	if size == 0 {
		rm.logger.Info("RemoverMember - Run - User not found")
		return errUser
	}

	id , err := valueobject.ValidateId(userId)
	if err != nil {
		rm.logger.Info("RemoverMember - Run - User Id is not valid")
		return err
	}

	errRemove := team.RemoveMember(valueobject.Id(id))

	if errRemove != nil {
		rm.logger.Info("RemoverMember - Run - User not found")
		return errRemove
	}

	return rm.teamUpdater.Run(ctx, team)
}