package repository

import (
	"context"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

type UserRepository interface {
	Find(ctx context.Context, userId valueobject.Id) (pkg.Optional[model.User], error)
	Save(ctx context.Context, user model.User) error
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Find(ctx context.Context, userId valueobject.Id) (pkg.Optional[model.User], error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(pkg.Optional[model.User]) , args.Error(1)
}

func (m *MockUserRepository) Save(ctx context.Context, user model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}