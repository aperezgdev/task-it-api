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

func NewTeam(title, description, owner string) (Team, error) {
	idVO, errId := valueobject.NewId()
	if errId != nil {
		return Team{}, errId
	}
	titleVO, errTitle := valueobject.NewTitle(title)
	if errTitle != nil {
		return Team{}, errTitle
	}
	descriptionVO, errDescription := valueobject.NewDescription(description)
	if errDescription != nil {
		return Team{}, errDescription
	}
	ownerVO, errOwner := valueobject.ValidateId(owner)
	if errOwner != nil {
		return Team{}, errOwner
	}

	return Team{
		Id: idVO,
		Title: titleVO,
		Description: descriptionVO,
		Members: []valueobject.Id{},
		Owner: ownerVO,
		CreatedAt: valueobject.NewCreatedAt(),
	}, nil
}

func (t *Team) AddMember(userId valueobject.Id) {
	t.Members = append(t.Members, userId)
}