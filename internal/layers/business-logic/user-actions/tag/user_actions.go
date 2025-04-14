package tag

import "medicine/internal/layers/business-logic/authorization"

type UserActions struct {
	ForceCreate       *ForceCreateUA
	UntagAllAndDelete *UntagAllAndDeleteUA
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UserActions {
	forceCreateUA := NewForceCreateUA(authorizer, simpleActions)
	untagAllAndDelete := NewUntagAllAndDeleteUA(authorizer, simpleActions)

	return &UserActions{
		ForceCreate:       forceCreateUA,
		UntagAllAndDelete: untagAllAndDelete,
	}
}
