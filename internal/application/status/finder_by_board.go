package status

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/board"
)

type StatusFinderByBoard struct {
	logger slog.Logger
	repository repository.StatusRepository
	boardFinder board.BoardFinder
}

func NewStatusFinderByBoard(logger slog.Logger, repository repository.StatusRepository, boardRepository repository.BoardRepository) StatusFinderByBoard {
	return StatusFinderByBoard{logger, repository, board.NewBoardFinder(logger, boardRepository)}
}

func (sf StatusFinderByBoard) Run(ctx context.Context, boardId string) ([]model.Status, error) {
	sf.logger.Info("StatusFinderByBoard - Run - Params: ", slog.Any("boardId", boardId))
	board, errBoard := sf.boardFinder.Run(ctx, boardId)
	if errBoard != nil {
		sf.logger.Info("StatusFinderByBoard - Run - Board not found")
		return []model.Status{}, errBoard
	}

	statusOptional, err := sf.repository.FindByBoard(ctx, board.Id)
	if err != nil {
		return []model.Status{}, err
	}

	if !statusOptional.IsPresent {
		return []model.Status{}, nil
	}

	return statusOptional.Value, nil
}