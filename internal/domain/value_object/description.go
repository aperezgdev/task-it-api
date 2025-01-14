package valueobject

import "github.com/aperezgdev/task-it-api/internal/domain/errors"

type Description string

func NewDescription(value string) (Description, error) {
	if len(value) > 240 {
		return "", errors.NewValidationError("Description", "must not exceed 240 characters")
	}
	return Description(value), nil
}