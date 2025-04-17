package tags_space

import (
	"medicine/internal/layers/business-logic/authorization"
)

type UserActions struct {
	GetByID          *GetByIDUA
	ListAllAvailable *ListAllAvailableUA
	Create           *CreateUA
	Delete           *DeleteUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	getByID := NewGetByIDUA(authorizer, simpleActions)
	listAllAvailableUser := NewListAllAvailableUA(simpleActions)
	create := NewCreateUA(authorizer, simpleActions)
	delete_ := NewDeleteUA(authorizer, simpleActions)

	return &UserActions{
		GetByID:          getByID,
		ListAllAvailable: listAllAvailableUser,
		Create:           create,
		Delete:           delete_,
	}
}
