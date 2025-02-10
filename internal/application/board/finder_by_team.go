package board

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/team"
)

type BoardFinderByTeam struct {
	logger slog.Logger
	repository repository.BoardRepository
	teamFinder team.TeamFinder
}

func NewBoardFinderByTeam(logger slog.Logger, repository repository.BoardRepository, teamRepository repository.TeamRepository) BoardFinderByTeam {
	return BoardFinderByTeam{logger, repository, team.NewTeamFinder(logger, teamRepository)}
}

func (bf BoardFinderByTeam) Run(ctx context.Context, teamId string) ([]model.Board, error) {
	bf.logger.Info("BoardFinderByTeam - Run - Params: ", slog.Any("teamId", teamId))
	team, errTeam := bf.teamFinder.Run(ctx, teamId)
	if errTeam != nil {
		bf.logger.Info("BoardFinderByTeam - Run - Team not found")
		return []model.Board{}, errTeam
	}

	boardOptional, err := bf.repository.FindByTeam(ctx, team.Id)
	if err != nil {
		return []model.Board{}, err
	}

	if !boardOptional.IsPresent {
		return []model.Board{}, nil
	}

	return boardOptional.Value, nil
}