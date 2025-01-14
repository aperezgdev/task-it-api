package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type Board struct {
	Id valueobject.Id
	Title valueobject.Title
	Description valueobject.Description
	OwnerId valueobject.Id
	CreatedAt valueobject.CreatedAt
}