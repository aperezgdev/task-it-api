package board

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type BoardFinder struct {
	logger slog.Logger
	repository repository.BoardRepository
}

func NewBoardFinder(logger slog.Logger, repository repository.BoardRepository) BoardFinder {
	return BoardFinder{logger, repository}
}

func (bf *BoardFinder) Run(ctx context.Context, idBoard string) (model.Board, error) {
	bf.logger.Info("BoardFinder - Run - Params: ", slog.Any("idBoard", idBoard))
	id, errId := valueobject.ValidateId(idBoard)
	if errId != nil {
		bf.logger.Info("BoardFinder - Run - Board id was invalid")
		return model.Board{}, errId
	}

	boardOptional, err := bf.repository.Find(ctx, id)
	if err != nil {
		return model.Board{}, errId
	}

	if !boardOptional.IsPresent {
		return model.Board{}, errors.ErrNotExist
	}

	return boardOptional.Value, nil

}