package tags_space

import (
	"medicine/internal/layers/business-logic/authorization"
)

type UserActions struct {
	GetByID *GetByIDUA
	Create  *CreateUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	getByID := NewGetByIDUA(authorizer, simpleActions)
	create := NewCreateUA(authorizer, simpleActions)

	return &UserActions{
		GetByID: getByID,
		Create:  create,
	}
}
