package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type Team struct {
	Id valueobject.Id
	Title valueobject.Title
	Description valueobject.Description
	Members []valueobject.Id
	Owner valueobject.Id
	CreatedAt valueobject.CreatedAt
}