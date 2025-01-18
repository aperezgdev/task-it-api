package model

import (
	valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"
)

type Status struct {
	Id valueobject.Id
	Title valueobject.Title
	BoardId valueobject.Id
	NextStatus []valueobject.Id
	PreviousStatus []valueobject.Id
	CreatedAt valueobject.CreatedAt
}

func NewStatus(title, boardId string, nextStatus, previousStatus []string) (Status, error) {
	idVO, errId := valueobject.NewId()
	if errId != nil {
		return Status{}, errId
	}

	tilteVO, errTitle := valueobject.NewTitle(title)
	if errTitle != nil {
		return Status{}, errTitle
	}

	boardIdVO, errBoardId := valueobject.ValidateId(boardId)
	if errBoardId != nil {

	}

	nextStatusVOs := make([]valueobject.Id, len(nextStatus))
	for _, v := range nextStatus {
		nextStatusVO, errNextStatus := valueobject.ValidateId(v)
		if errNextStatus != nil {
			return Status{}, errNextStatus
		}
		nextStatusVOs = append(nextStatusVOs, nextStatusVO)
	}

	previousStatusVOs := make([]valueobject.Id, len(previousStatus))
	for _, v := range previousStatus {
		previousStatusVO, errPreviousStatus := valueobject.ValidateId(v)
		if errPreviousStatus != nil {
			return Status{}, errPreviousStatus
		}
		previousStatusVOs = append(previousStatusVOs, previousStatusVO)
	}

	return Status{
		Id: idVO,
		Title: tilteVO,
		BoardId: boardIdVO,
		NextStatus: nextStatusVOs,
		PreviousStatus: previousStatusVOs,
		CreatedAt: valueobject.NewCreatedAt(),
	}, nil
}