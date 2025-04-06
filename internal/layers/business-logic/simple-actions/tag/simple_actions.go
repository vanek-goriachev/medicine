package tag

type SimpleActions struct {
	idGenerator EntityIDGenerator
	tagFactory  TagFactory

	atomicActions AtomicActions
}

func NewSimpleActions(
	idGenerator EntityIDGenerator,
	tagFactory TagFactory,
	atomicActions AtomicActions,
) *SimpleActions {
	return &SimpleActions{
		idGenerator:   idGenerator,
		tagFactory:    tagFactory,
		atomicActions: atomicActions,
	}
}
