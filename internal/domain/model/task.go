package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type Task struct {
	Id valueobject.Id
	Title valueobject.Title
	Description valueobject.Description
	Creator valueobject.Id
	Asigned valueobject.Id
	StatusId valueobject.Id
	CreatedAt valueobject.CreatedAt
}

func NewTask(title, description, creator, asigned, statusId string) (Task, error) {
	idVO, errId := valueobject.NewId()
	if errId != nil {
		return Task{}, errId
	}

	titleVO, errTitle := valueobject.NewTitle(title)
	if errTitle != nil {
		return Task{}, errTitle	
	}

	descriptionVO, errDescription := valueobject.NewDescription(description)
	if errDescription != nil {
		return Task{}, errDescription
	}

	creatorVO, errCreator := valueobject.ValidateId(creator)
	if errCreator != nil {
		return Task{}, errCreator
	}

	asignedVO, errAsigned := valueobject.ValidateId(asigned)
	if errAsigned != nil {
		return Task{}, errAsigned
	}

	statusIdVO, errStatusId := valueobject.ValidateId(statusId)
	if errStatusId != nil {
		return Task{}, errStatusId
	}

	return Task{
		Id: idVO,
		Title: titleVO,
		Description: descriptionVO,
		Creator: creatorVO,
		Asigned: asignedVO,
		StatusId: statusIdVO,
		CreatedAt: valueobject.NewCreatedAt(),
	}, nil
}