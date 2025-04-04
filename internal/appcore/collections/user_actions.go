package collections

import (
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
)

type UserActions struct {
	TagsSpace *tagsSpaceUA.UserActions
}

func NewUserActions(simpleActions *SimpleActions) *UserActions {
	var c UserActions

	c.TagsSpace = tagsSpaceUA.NewUserActions(
		simpleActions.tagsSpace,
	)

	return &c
}
