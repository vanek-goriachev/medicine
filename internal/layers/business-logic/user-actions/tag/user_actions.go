package tag

import "medicine/internal/layers/business-logic/authorization"

type UserActions struct {
	ForceCreate *ForceCreateUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	forceCreateUA := NewForceCreateUA(authorizer, simpleActions)

	return &UserActions{
		ForceCreate: forceCreateUA,
	}
}
