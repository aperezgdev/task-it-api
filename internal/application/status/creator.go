package status

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/board"
)

type StatusCreator struct {
	logger slog.Logger
	repository repository.StatusRepository
	boardFinder board.BoardFinder
}

func NewStatusCreator(logger slog.Logger, repository repository.StatusRepository, boardRepository repository.BoardRepository) StatusCreator {
	return StatusCreator{logger, repository, board.NewBoardFinder(logger, boardRepository)}
}

func (sc *StatusCreator) Run(ctx context.Context, title, idBoard string, nextStatus, previousStatus []string) error {
	sc.logger.Info("StatusCreator - Run - Params: ", slog.Any("title", title), slog.Any("idBoard", idBoard), slog.Any("nextStatus", nextStatus), slog.Any("previousStatus", previousStatus))
	_, errBoard := sc.boardFinder.Run(ctx, idBoard)

	if errBoard != nil {
		sc.logger.Info("StatusCreator - Run - Board not found")
		return errBoard
	}

	status, err := model.NewStatus(title, idBoard, nextStatus, previousStatus)
	if err != nil {
		sc.logger.Info("StatusCreator - Run - Error trying to create status")
		return err
	}

	return sc.repository.Save(ctx, status)
}