package board

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/board"
)

type BoardRemover struct {
	logger slog.Logger
	repository repository.BoardRepository
	finder board.BoardFinder
}

func NewBoardRemover(logger slog.Logger, repository repository.BoardRepository) BoardRemover {
	return BoardRemover{
		logger: logger,
		repository: repository,
		finder: board.NewBoardFinder(logger, repository),
	}
}

func (br *BoardRemover) Run(ctx context.Context, id string) error {	
	board, err := br.finder.Run(ctx, id)
	if err != nil {
		br.logger.Info("BoardRemover - Run - Error trying to find board")
		return err
	}

	err = br.repository.Delete(ctx, board.Id)
	if err != nil {
		br.logger.Info("BoardRemover - Run - Error trying to delete board")
		return err
	}

	return nil
}