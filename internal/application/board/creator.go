package board

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/team"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/user"
)

type BoardCreator struct {
	logger slog.Logger
	boardRepository repository.BoardRepository
	userFinder user.UserFinder
	teamFinder team.TeamFinder
}

func NewBoardCreator(logger slog.Logger, boardRepository repository.BoardRepository, userRepository repository.UserRepository, teamRepository repository.TeamRepository) BoardCreator {
	return BoardCreator{
		logger: logger, 
		boardRepository: boardRepository, 
		userFinder:  user.NewUserFinder(logger, userRepository),
		teamFinder:  team.NewTeamFinder(logger, teamRepository),
	}
}

func (bc BoardCreator) Run(ctx context.Context, title, description, owner, teamId string) error {
	bc.logger.Info(
		"BoardCreator - Run - Params: ", 
		slog.Any("title", title), 
		slog.Any("description", description),
		slog.Any("owner", owner),
		slog.Any("teamId", teamId),
	)

	board, err := model.NewBoard(title, description, owner, teamId)
	if err != nil {
		bc.logger.Info("BoardCreator - Run - Error trying to create board")
		return err
	}

	_, errTeam := bc.teamFinder.Run(ctx, teamId)
	if errTeam != nil {
		bc.logger.Info("BoardCreator - Run - Team not found")
		return errTeam
	}

	_, errUser := bc.userFinder.Run(ctx, owner)
	if errUser != nil {
		bc.logger.Info("BoardCreator - Run - User not found")
		return errUser
	}

	return bc.boardRepository.Save(ctx, board)
}