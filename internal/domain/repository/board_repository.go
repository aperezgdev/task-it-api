package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

type BoardRepository interface {
	Find(ctx context.Context, boardId valueobject.Id) (pkg.Optional[model.Board], error)
	FindByTeam(ctx context.Context, teamId valueobject.Id) (pkg.Optional[[]model.Board], error)
	Save(ctx context.Context, board model.Board) error
	Delete(ctx context.Context, boardId valueobject.Id) error
	Update(ctx context.Context, board model.Board) error
}

func (m *MockBoardRepository) FindByTeam(ctx context.Context, id valueobject.Id) (pkg.Optional[[]model.Board], error) {
	args := m.Called(ctx, id)
	return args.Get(0).(pkg.Optional[[]model.Board]), args.Error(1)
}

type MockBoardRepository struct {
	mock.Mock
}

func (m *MockBoardRepository) Find(ctx context.Context, boardId valueobject.Id) (pkg.Optional[model.Board], error) {
	args := m.Called(ctx, boardId)
	return args.Get(0).(pkg.Optional[model.Board]), args.Error(1)
}

func (m *MockBoardRepository) Save(ctx context.Context, board model.Board) error {
	args := m.Called(ctx, board)
	return args.Error(0)
}

func (m *MockBoardRepository) Delete(ctx context.Context, boardId valueobject.Id) error {
	args := m.Called(ctx, boardId)
	return args.Error(0)
}

func (m *MockBoardRepository) Update(ctx context.Context, board model.Board) error {
	args := m.Called(ctx, board)
	return args.Error(0)
}