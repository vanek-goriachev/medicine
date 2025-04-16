package visit_record

import "medicine/internal/layers/business-logic/authorization"

type UserActions struct {
	Create *CreateUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	createUA := NewCreateUA(authorizer, simpleActions)

	return &UserActions{
		Create: createUA,
	}
}
