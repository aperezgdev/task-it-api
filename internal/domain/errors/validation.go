package errors

import "fmt"

type ValidationError struct {
    Field string
    Msg   string
}

func NewValidationError(field, msg string) ValidationError {
    return ValidationError{
        Field: field,
        Msg:   msg,
    }
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Msg)
}