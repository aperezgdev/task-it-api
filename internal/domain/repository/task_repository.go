package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

type TaskRepository interface {
	Find(ctx context.Context, taskId valueobject.Id) (pkg.Optional[model.Task], error)
	Save(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, taskId valueobject.Id) error
	Update(ctx context.Context, task model.Task) error
}

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Find(ctx context.Context, taskId valueobject.Id) (pkg.Optional[model.Task], error) {
	args := m.Called(ctx, taskId)
	return args.Get(0).(pkg.Optional[model.Task]), args.Error(1)
}

func (m *MockTaskRepository) Save(ctx context.Context, task model.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(ctx context.Context, taskId valueobject.Id) error {
	args := m.Called(ctx, taskId)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(ctx context.Context, task model.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}