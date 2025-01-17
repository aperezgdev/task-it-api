package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/stretchr/testify/mock"
)

type StatusRepository interface {
	Save(ctx context.Context, status model.Status) error
	Delete(ctx context.Context, statusId valueobject.Id) error
	Update(ctx context.Context, status model.Status) error
}

type MockStatusRepository struct {
	mock.Mock
}

func (m *MockStatusRepository) Save(ctx context.Context, status model.Status) error {
	args := m.Called(ctx, status)
	return args.Error(0)
}

func (m *MockStatusRepository) Delete(ctx context.Context, statusId valueobject.Id) error {
	args := m.Called(ctx, statusId)
	return args.Error(0)
}

func (m *MockStatusRepository) Update(ctx context.Context, status model.Status) error {
	args := m.Called(ctx, status)
	return args.Error(0)
}