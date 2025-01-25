package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/postgresql/sqlc"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BoardRepository struct {
	queries *sqlc.Queries
	logger  slog.Logger
}

func NewBoardRepository(logger slog.Logger, queries *sqlc.Queries) BoardRepository {
	return BoardRepository{queries, logger}
}

func (br BoardRepository) Find(ctx context.Context, boardId valueobject.Id) (pkg.Optional[model.Board], error) {
	board, err := br.queries.FindBoard(ctx, uuid.UUID(boardId))
	if err != nil {
		br.logger.Error("BoardRepository - Find - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return pkg.EmptyOptional[model.Board](), err
	}

	return pkg.NewOptional(model.Board{
		Id: valueobject.Id(board.ID),
		Title: valueobject.Title(board.Title),
		Description: valueobject.Description(board.Description),
		Owner: valueobject.Id(board.Owner),
		Team: valueobject.Id(board.TeamID),
		CreatedAt: valueobject.CreatedAt(board.CreatedAt.Time),
	}), nil
}

func (br BoardRepository) Save(ctx context.Context, board model.Board) error {
	_, err := br.queries.SaveBoard(ctx, sqlc.SaveBoardParams{
		ID:          uuid.UUID(board.Id),
		Title:       string(board.Title),
		Description: string(board.Description),
		Owner:       uuid.UUID(board.Owner),
		TeamID:      uuid.UUID(board.Team),
		CreatedAt:   pgtype.Timestamp{Time: time.Time(board.CreatedAt)},
	})
	if err != nil {
		br.logger.Error("BoardRepository - Save - Error has ocurred trying to save data to repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (br BoardRepository) Delete(ctx context.Context, boardId valueobject.Id) error {
	err := br.queries.DeleteBoard(ctx, uuid.UUID(boardId))
	if err != nil {
		br.logger.Error("BoardRepository - Delete - Error has ocurred trying to delete data from repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (br BoardRepository) Update(ctx context.Context, board model.Board) error {
	err := br.queries.UpdateBoard(ctx, sqlc.UpdateBoardParams{
		ID:          uuid.UUID(board.Id),
		Title:       string(board.Title),
		Description: string(board.Description),
		Owner:       uuid.UUID(board.Owner),
		TeamID:      uuid.UUID(board.Team),
		CreatedAt:   pgtype.Timestamp{Time: time.Time(board.CreatedAt)},
	})
	if err != nil {
		br.logger.Error("BoardRepository - Update - Error has ocurred trying to update data from repository", slog.Any("error", err))
		return err
	}

	return nil
}