package collections

import (
	"medicine/internal/layers/business-logic/authorization"
	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
)

type UserActions struct {
	authorizer authorization.Authorizer

	Tag       *tagUA.UserActions
	TagsSpace *tagsSpaceUA.UserActions
}

func NewUserActions(
	authorizer authorization.Authorizer,
	simpleActions *SimpleActions,
) *UserActions {
	var c UserActions
	c.authorizer = authorizer

	c.Tag = tagUA.NewUserActions(
		c.authorizer,
		simpleActions.tag,
	)
	c.TagsSpace = tagsSpaceUA.NewUserActions(
		c.authorizer,
		simpleActions.tagsSpace,
	)

	return &c
}
