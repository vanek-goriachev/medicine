package tags_space

type SimpleActions struct {
	idGenerator      EntityIDGenerator
	tagsSpaceFactory TagsSpaceFactory
	atomicActions    AtomicActions
}

func NewSimpleActions(
	idGenerator EntityIDGenerator,
	tagsSpaceFactory TagsSpaceFactory,
	atomicActions AtomicActions,
) *SimpleActions {
	return &SimpleActions{
		idGenerator:      idGenerator,
		tagsSpaceFactory: tagsSpaceFactory,
		atomicActions:    atomicActions,
	}
}
