package valueobject

import "github.com/aperezgdev/task-it-api/internal/domain/errors"

type Title string

func NewTitle(value string) (Title, error) {
	if len(value) == 0 || len(value) > 40 {
		return "", errors.NewValidationError("Title", "must not be empty or exceed 24 characters")
	}
	return Title(value), nil
}