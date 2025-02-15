package local

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
)

type TeamRepository struct {
	logger slog.Logger
	data map[valueobject.Id]model.Team
}

func NewTeamRepository(logger slog.Logger) TeamRepository {
	return TeamRepository{logger, make(map[valueobject.Id]model.Team)}
}

func (tr TeamRepository) Find(ctx context.Context, teamId valueobject.Id) (pkg.Optional[model.Team], error) {
	tr.logger.Info("TeamRepository - Find - Params: ", slog.Any("teamId", teamId))
	team, ok := tr.data[teamId]
	if !ok {
		return pkg.EmptyOptional[model.Team](), nil
	}

	return pkg.NewOptional(team), nil
}

func (tr TeamRepository) FindByMember(ctx context.Context, memberId valueobject.Id) (pkg.Optional[model.Team], error) {
	tr.logger.Info("TeamRepository - FindByMember - Params: ", slog.Any("memberId", memberId))
	var team model.Team
	for _, t := range tr.data {
		for _, member := range t.Members {
			if member == memberId {
				team = t
				break
			}
		}
	}
	var emptyId valueobject.Id
	if team.Id == emptyId {
		return pkg.EmptyOptional[model.Team](), nil
	}

	return pkg.NewOptional(team), nil
}

func (tr TeamRepository) Save(ctx context.Context, team model.Team) error {
	tr.logger.Info("TeamRepository - Save - Params: ", slog.Any("team", team))
	tr.data[team.Id] = team
	return nil
}

func (tr TeamRepository) Delete(ctx context.Context, teamId valueobject.Id) error {
	tr.logger.Info("TeamRepository - Delete - Params: ", slog.Any("teamId", teamId))
	delete(tr.data, teamId)
	return nil
}

func (tr TeamRepository) Update(ctx context.Context, team model.Team) error {
	tr.logger.Info("TeamRepository - Update - Params: ", slog.Any("team", team))
	tr.data[team.Id] = team
	return nil
}