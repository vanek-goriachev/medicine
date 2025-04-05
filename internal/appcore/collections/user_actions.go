package collections

import (
	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
)

type UserActions struct {
	authorizer authorization.Authorizer

	TagsSpace *tagsSpaceUA.UserActions
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions *SimpleActions,
) *UserActions {
	var c UserActions
	c.authorizer = authorizer

	c.TagsSpace = tagsSpaceUA.NewUserActions(
		c.authorizer,
		simpleActions.tagsSpace,
	)

	return &c
}
