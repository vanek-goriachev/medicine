package tags_space

type SimpleActions struct {
	idGenerator            EntityIDGenerator
	tagsSpaceFactory       TagsSpaceFactory
	tagAtomicActions       TagAtomicActions
	tagsSpaceAtomicActions TagsSpaceAtomicActions
}

func NewSimpleActions(
	idGenerator EntityIDGenerator,
	tagsSpaceFactory TagsSpaceFactory,
	tagAtomicActions TagAtomicActions,
	tagsSpaceAtomicActions TagsSpaceAtomicActions,
) *SimpleActions {
	return &SimpleActions{
		idGenerator:            idGenerator,
		tagsSpaceFactory:       tagsSpaceFactory,
		tagAtomicActions:       tagAtomicActions,
		tagsSpaceAtomicActions: tagsSpaceAtomicActions,
	}
}
