package tags_space

import entityID "medicine/pkg/entity-id"

type SimpleActions struct {
	idGenerator      entityID.Generator
	tagsSpaceFactory TagsSpaceFactory
	atomicActions    AtomicActions
}

func NewSimpleActions(
	idGenerator entityID.Generator,
	tagsSpaceFactory TagsSpaceFactory,
	atomicActions AtomicActions,
) *SimpleActions {
	return &SimpleActions{
		idGenerator:      idGenerator,
		tagsSpaceFactory: tagsSpaceFactory,
		atomicActions:    atomicActions,
	}
}
