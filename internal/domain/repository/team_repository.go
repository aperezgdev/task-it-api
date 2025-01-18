package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

type TeamRepository interface {
	Find(ctx context.Context, idTeam valueobject.Id) (pkg.Optional[model.Team], error)
	Save(ctx context.Context, team model.Team) error
	Delete(ctx context.Context, idTeam valueobject.Id) error
	Update(ctx context.Context, team model.Team) error
}

type MockTeamRepository struct {
	mock.Mock
}

func (m *MockTeamRepository) Find(ctx context.Context, idTeam valueobject.Id) (pkg.Optional[model.Team], error) {
	args := m.Called(ctx, idTeam)
	return args.Get(0).(pkg.Optional[model.Team]), args.Error(1)
}

func (m *MockTeamRepository) Save(ctx context.Context, team model.Team) error {
	args := m.Called(ctx, team)
	return args.Error(0)
}

func (m *MockTeamRepository) Delete(ctx context.Context, idTeam valueobject.Id) error {
	args := m.Called(ctx, idTeam)
	return args.Error(0)
}

func (m *MockTeamRepository) Update(ctx context.Context, team model.Team) error {
	args := m.Called(ctx, team)
	return args.Error(0)
}