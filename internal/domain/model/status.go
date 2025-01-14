package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type Status struct {
	Id valueobject.Id
	Title valueobject.Title
	BoardId valueobject.Id
	NextStatus []valueobject.Id
	PreviousStatus []valueobject.Id
	CreatedAt valueobject.CreatedAt
}