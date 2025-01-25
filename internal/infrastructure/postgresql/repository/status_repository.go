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

type StatusRepository struct {
	queries *sqlc.Queries
	logger  slog.Logger
}

func NewStatusRepository(logger slog.Logger, queries *sqlc.Queries) StatusRepository {
	return StatusRepository{queries, logger}
}

func (sr StatusRepository) Find(ctx context.Context, statusId valueobject.Id) (pkg.Optional[model.Status], error) {
	status, err := sr.queries.FindStatus(ctx, uuid.UUID(statusId))
	if err != nil {
		sr.logger.Error("StatusRepository - Find - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return pkg.EmptyOptional[model.Status](), err
	}

	var nextStatus []valueobject.Id
	if ns, ok := status.NextStatus.([]valueobject.Id); ok {
		nextStatus = ns
	} else {
		nextStatus = make([]valueobject.Id, 0)
	}

	var previousStatus []valueobject.Id
	if ps, ok := status.PreviousStatus.([]valueobject.Id); ok {
		previousStatus = ps
	} else {
		previousStatus = make([]valueobject.Id, 0)
	}

	return pkg.NewOptional(model.Status{
		Id:           valueobject.Id(status.ID),
		Title:        valueobject.Title(status.Title),
		BoardId:      valueobject.Id(status.BoardID),
		NextStatus:   nextStatus,
		PreviousStatus: previousStatus,
		CreatedAt: valueobject.CreatedAt(status.CreatedAt.Time),
	}), nil
}

func (sr StatusRepository) Save(ctx context.Context, status model.Status) error {
	nextStatusUUIDs := make([]uuid.UUID, len(status.NextStatus))
	for i, id := range status.NextStatus {
		nextStatusUUIDs[i] = uuid.UUID(id)
	}
	
	previousStatusUUIDs := make([]uuid.UUID, len(status.PreviousStatus))
	for i, id := range status.PreviousStatus {
		previousStatusUUIDs[i] = uuid.UUID(id)
	}

	_, err := sr.queries.SaveStatus(ctx, sqlc.SaveStatusParams{
		ID:             uuid.UUID(status.Id),
		Title:          string(status.Title),
		BoardID:        uuid.UUID(status.BoardId),
		CreatedAt:      pgtype.Timestamp{Time: time.Time(status.CreatedAt)},
	})
	if err != nil {
		sr.logger.Error("StatusRepository - Save - Error has ocurred trying to save data to repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (sr StatusRepository) Delete(ctx context.Context, statusId valueobject.Id) error {
	err := sr.queries.DeleteStatus(ctx, uuid.UUID(statusId))
	if err != nil {
		sr.logger.Error("StatusRepository - Delete - Error has ocurred trying to delete data from repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (sr StatusRepository) Update(ctx context.Context, status model.Status) error {
	err := sr.queries.UpdateStatus(ctx, sqlc.UpdateStatusParams{
		ID:             uuid.UUID(status.Id),
		Title:          string(status.Title),
		BoardID:        uuid.UUID(status.BoardId),
		CreatedAt:      pgtype.Timestamp{Time: time.Time(status.CreatedAt)},
	})
	if err != nil {
		sr.logger.Error("StatusRepository - Update - Error has ocurred trying to update data from repository", slog.Any("error", err))
		return err
	}

	return nil
}