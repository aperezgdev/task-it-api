package local

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
)

type BoardRepository struct {
	logger slog.Logger
	data map[valueobject.Id]model.Board
}

func NewBoardRepository(logger slog.Logger) BoardRepository {
	return BoardRepository{logger, make(map[valueobject.Id]model.Board)}
}

func (br BoardRepository) Find(ctx context.Context, boardId valueobject.Id) (pkg.Optional[model.Board], error) {
	br.logger.Info("BoardRepository - Find - Params: ", slog.Any("boardId", boardId))
	board, ok := br.data[boardId]
	if !ok {
		return pkg.EmptyOptional[model.Board](), nil
	}

	return pkg.NewOptional(board), nil
}

func (br BoardRepository) FindByTeam(ctx context.Context, teamId valueobject.Id) (pkg.Optional[[]model.Board], error) {
	br.logger.Info("BoardRepository - FindByTeam - Params: ", slog.Any("teamId", teamId))
	var boards []model.Board
	for _, board := range br.data {
		if board.Team == teamId {
			boards = append(boards, board)
		}
	}
	
	if len(boards) == 0 {
		return pkg.EmptyOptional[[]model.Board](), nil
	}

	return pkg.NewOptional(boards), nil
}

func (br BoardRepository) Save(ctx context.Context, board model.Board) error {
	br.logger.Info("BoardRepository - Save - Params: ", slog.Any("board", board))
	br.data[board.Id] = board
	return nil
}

func (br BoardRepository) Delete(ctx context.Context, boardId valueobject.Id) error {
	br.logger.Info("BoardRepository - Delete - Params: ", slog.Any("boardId", boardId))
	delete(br.data, boardId)
	return nil
}

func (br BoardRepository) Update(ctx context.Context, board model.Board) error {
	br.logger.Info("BoardRepository - Update - Params: ", slog.Any("board", board))
	br.data[board.Id] = board
	return nil
}