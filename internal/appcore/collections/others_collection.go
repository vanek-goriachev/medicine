package collections

import entityID "medicine/pkg/entity-id"

type Others struct {
	entityIDGenerator entityID.Generator
}

func NewOthers() *Others {
	var c Others

	c.entityIDGenerator = entityID.NewGenerator()

	return &c
}
