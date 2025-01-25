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

type TeamRepository struct {
	queries *sqlc.Queries
	logger  slog.Logger
}

func NewTeamRepository(logger slog.Logger, queries *sqlc.Queries) TeamRepository {
	return TeamRepository{queries, logger}
}

func (tr TeamRepository) Find(ctx context.Context, idTeam valueobject.Id) (pkg.Optional[model.Team], error) {
	team, err := tr.queries.FindTeam(ctx, uuid.UUID(idTeam))
	if err != nil {
		tr.logger.Error("TeamRepository - Find - Error has ocurred trying to retrieve data from repository", slog.Any("error", err))
		return pkg.EmptyOptional[model.Team](), err
	}

	var members []valueobject.Id
	if ms, ok := team.Members.([]valueobject.Id); ok {
		members = ms
	} else {
		members = make([]valueobject.Id, 0)
	}

	return pkg.NewOptional(model.Team{
		Id: valueobject.Id(team.ID),
		Title: valueobject.Title(team.Title),
		Description: valueobject.Description(team.Description),
		Members: members,
		Owner: valueobject.Id(team.Owner),
		CreatedAt: valueobject.CreatedAt(team.CreatedAt.Time),
	}), nil
}

func (tr TeamRepository) Save(ctx context.Context, team model.Team) error {
	_, err := tr.queries.SaveTeam(ctx, sqlc.SaveTeamParams{
		ID:          uuid.UUID(team.Id),
		Title:       string(team.Title),
		Description: string(team.Description),
		Owner:       uuid.UUID(team.Owner),
		CreatedAt:   pgtype.Timestamp{Time: time.Time(team.CreatedAt)},
	})
	if err != nil {
		tr.logger.Error("TeamRepository - Save - Error has ocurred trying to save data to repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (tr TeamRepository) Delete(ctx context.Context, idTeam valueobject.Id) error {
	err := tr.queries.DeleteTeam(ctx, uuid.UUID(idTeam))
	if err != nil {
		tr.logger.Error("TeamRepository - Delete - Error has ocurred trying to delete data from repository", slog.Any("error", err))
		return err
	}

	return nil
}

func (tr TeamRepository) Update(ctx context.Context, team model.Team) error {
	err := tr.queries.UpdateTeam(ctx, sqlc.UpdateTeamParams{
		ID:          uuid.UUID(team.Id),
		Title:       string(team.Title),
		Description: string(team.Description),
		Owner:       uuid.UUID(team.Owner),
		CreatedAt:   pgtype.Timestamp{Time: time.Time(team.CreatedAt)},
	})
	if err != nil {
		tr.logger.Error("TeamRepository - Update - Error has ocurred trying to update data from repository", slog.Any("error", err))
		return err
	}

	return nil
}	