package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type Board struct {
	Id valueobject.Id
	Title valueobject.Title
	Description valueobject.Description
	Owner valueobject.Id
	Team valueobject.Id
	CreatedAt valueobject.CreatedAt
}

func NewBoard(title, description, owner, team string) (Board, error) {
	idVO, errId := valueobject.NewId()
	if errId != nil {
		return Board{}, errId
	}

	titleVO, errTitle := valueobject.NewTitle(title)
	if errTitle != nil {
		return Board{}, errTitle
	}

	descriptionVO, errDescription := valueobject.NewDescription(description)
	if errDescription != nil {
		return Board{}, errDescription
	}

	ownerVO, errOwner := valueobject.ValidateId(owner)
	if errOwner != nil {
		return Board{}, errOwner
	}

	teamVO, errTeam := valueobject.ValidateId(team)
	if errTeam != nil {
		return Board{}, errTeam
	}

	return Board{
		Id: idVO,
		Title: titleVO,
		Description: descriptionVO,
		Owner: ownerVO,
		Team: teamVO,
		CreatedAt: valueobject.NewCreatedAt(),
	}, nil 
}	