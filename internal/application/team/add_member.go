package team

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/team"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/user"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type TeamAddMember struct {
	logger slog.Logger
	teamFinder team.TeamFinder
	teamUpdater team.TeamUpdater
	userFinder user.UserFinder
}

func NewTeamAddMember(logger slog.Logger, teamRepository repository.TeamRepository, userRepository repository.UserRepository) TeamAddMember {
	return TeamAddMember{logger,  team.NewTeamFinder(logger, teamRepository), team.NewTeamUpdater(logger, teamRepository), user.NewUserFinder(logger, userRepository)}
}

func (tam TeamAddMember) Run(ctx context.Context, idTeam, userId string) error {
	tam.logger.Info("TeamAddMember - Run - Params: ", slog.Any("idTeam", idTeam), slog.Any("userId", userId))
	_, errUser := tam.userFinder.Run(ctx, userId)
	if errUser != nil {
		tam.logger.Info("TeamAddMember - Run - User not found")
		return errUser
	}

	team, err := tam.teamFinder.Run(ctx, idTeam)
	if err != nil {
		tam.logger.Info("TeamAddMember - Run - Team not found")
		return err
	}

	userIdBytes := []byte(userId)

	var userIdArray [16]byte
	copy(userIdArray[:], userIdBytes)

	team.AddMember(valueobject.Id(userIdArray))

	teamToUpdate := model.Team{
		Id: team.Id,
		Members: team.Members,
	}

	return tam.teamUpdater.Run(ctx, teamToUpdate)
}