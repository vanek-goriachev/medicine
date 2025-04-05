package tags_space

import "medicine/internal/layers/business-logic/authorization"

type UserActions struct {
	Create *CreateUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	create := NewCreateUA(authorizer, simpleActions)

	return &UserActions{
		Create: create,
	}
}
