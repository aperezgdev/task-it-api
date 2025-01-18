package valueobject

import (
	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/google/uuid"
)

type Id uuid.UUID

func NewId() (Id, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return Id{}, err
	}
	return Id(id), nil
}

func ValidateId(value string) (Id , error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return Id{}, errors.NewValidationError("Id", "Id is not valid")
	}
	return Id(id), nil
}