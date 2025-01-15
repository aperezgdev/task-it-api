package valueobject

import (
	"net/mail"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
)

type Email string

func NewEmail(value string) (Email, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return "", errors.NewValidationError("Email", "Email format is not valid")
	}
	return Email(value), nil
}