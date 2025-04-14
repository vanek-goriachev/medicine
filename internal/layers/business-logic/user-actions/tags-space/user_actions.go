package tags_space

import (
	"medicine/internal/layers/business-logic/authorization"
)

type UserActions struct {
	GetByID    *GetByIDUA
	ListByUser *ListByUserUA
	Create     *CreateUA
	Delete     *DeleteUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	getByID := NewGetByIDUA(authorizer, simpleActions)
	listByUser := NewListByUserUA(authorizer, simpleActions)
	create := NewCreateUA(authorizer, simpleActions)
	delete_ := NewDeleteUA(authorizer, simpleActions)

	return &UserActions{
		GetByID:    getByID,
		ListByUser: listByUser,
		Create:     create,
		Delete:     delete_,
	}
}
