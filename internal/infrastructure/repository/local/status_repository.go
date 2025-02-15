package local

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
)

type StatusRepository struct {
	logger slog.Logger
	data map[valueobject.Id]model.Status
}

func NewStatusRepository(logger slog.Logger) StatusRepository {
	return StatusRepository{logger, make(map[valueobject.Id]model.Status)}
}

func (sr StatusRepository) Find(ctx context.Context, statusId valueobject.Id) (pkg.Optional[model.Status], error) {
	sr.logger.Info("StatusRepository - Find - Params: ", slog.Any("statusId", statusId))
	status, ok := sr.data[statusId]
	if !ok {
		return pkg.EmptyOptional[model.Status](), nil
	}

	return pkg.NewOptional(status), nil
}

func (sr StatusRepository) FindByBoard(ctx context.Context, boardId valueobject.Id) (pkg.Optional[[]model.Status], error) {
	sr.logger.Info("StatusRepository - FindByBoard - Params: ", slog.Any("boardId", boardId))
	var statuses []model.Status
	for _, status := range sr.data {
		if status.BoardId == boardId {
			statuses = append(statuses, status)
		}
	}
	
	if len(statuses) == 0 {
		return pkg.EmptyOptional[[]model.Status](), nil
	}

	return pkg.NewOptional(statuses), nil
}

func (sr StatusRepository) Save(ctx context.Context, status model.Status) error {
	sr.logger.Info("StatusRepository - Save - Params: ", slog.Any("status", status))
	sr.data[status.Id] = status
	return nil
}

func (sr StatusRepository) Delete(ctx context.Context, statusId valueobject.Id) error {
	sr.logger.Info("StatusRepository - Delete - Params: ", slog.Any("statusId", statusId))
	delete(sr.data, statusId)
	return nil
}

func (sr StatusRepository) Update(ctx context.Context, status model.Status) error {
	sr.logger.Info("StatusRepository - Update - Params: ", slog.Any("status", status))
	sr.data[status.Id] = status
	return nil
}