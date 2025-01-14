package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type Task struct {
	Id valueobject.Id
	Title valueobject.Title
	Description valueobject.Description
	UserId valueobject.Id
	StatusId valueobject.Id
	CreatedAt valueobject.CreatedAt
}