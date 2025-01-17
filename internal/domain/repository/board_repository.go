package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/stretchr/testify/mock"
)

type BoardRepository interface {
	Save(ctx context.Context, board model.Board) error
	Delete(ctx context.Context, boardId valueobject.Id) error
	Update(ctx context.Context, board model.Board) error
}

type MockBoardRepository struct {
	mock.Mock
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