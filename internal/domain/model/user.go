package model

import valueobject "github.com/aperezgdev/task-it-api/internal/domain/value_object"

type User struct {
	Id valueobject.Id
	Email valueobject.Email
	CreatedAt valueobject.CreatedAt
}

func NewUser(email string) (User, error) {
	idVO, errId := valueobject.NewId()
	if errId != nil {
		return User{}, errId
	}

	emailVO, errEmail := valueobject.NewEmail(email)
	if errEmail != nil {
		return User{}, errEmail
	}

	return User{
		Id: idVO,
		Email: emailVO,
		CreatedAt: valueobject.NewCreatedAt(),
	}, nil	
}