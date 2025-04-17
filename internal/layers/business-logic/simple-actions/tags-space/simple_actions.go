package tags_space

import "medicine/internal/layers/business-logic/authorization"

type SimpleActions struct {
	authorizer             authorization.Authorizer
	idGenerator            EntityIDGenerator
	tagsSpaceFactory       TagsSpaceFactory
	tagAtomicActions       TagAtomicActions
	tagsSpaceAtomicActions TagsSpaceAtomicActions
}

func NewSimpleActions(
	authorizer authorization.Authorizer,
	idGenerator EntityIDGenerator,
	tagsSpaceFactory TagsSpaceFactory,
	tagAtomicActions TagAtomicActions,
	tagsSpaceAtomicActions TagsSpaceAtomicActions,
) *SimpleActions {
	return &SimpleActions{
		authorizer:             authorizer,
		idGenerator:            idGenerator,
		tagsSpaceFactory:       tagsSpaceFactory,
		tagAtomicActions:       tagAtomicActions,
		tagsSpaceAtomicActions: tagsSpaceAtomicActions,
	}
}
