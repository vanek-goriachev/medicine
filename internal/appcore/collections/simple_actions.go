package collections

import (
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
)

type SimpleActions struct {
	tagsSpace *tagsSpaceSA.SimpleActions
}

func NewSimpleActions(
	others *Others,
	gateways *Gateways, // Using gateway instead of AtomicAction because of the same interface
	factories *Factories,
) *SimpleActions {
	var c SimpleActions

	c.tagsSpace = tagsSpaceSA.NewSimpleActions(
		others.entityIDGenerator,
		factories.tagsSpace,
		gateways.tagsSpace, // Using gateway instead of AtomicAction because of the same interface
	)

	return &c
}
