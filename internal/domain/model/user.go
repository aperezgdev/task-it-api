package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type User struct {
	Id valueobject.Id
	Email valueobject.Email
	CreatedAt valueobject.CreatedAt
}