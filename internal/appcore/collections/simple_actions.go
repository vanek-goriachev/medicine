package collections

import (
	tagSA "medicine/internal/layers/business-logic/simple-actions/tag"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
)

type SimpleActions struct {
	tag       *tagSA.SimpleActions
	tagsSpace *tagsSpaceSA.SimpleActions
}

func NewSimpleActions(
	others *Others,
	gateways *Gateways, // Using gateway instead of AtomicAction because of the same interface
	factories *Factories,
) *SimpleActions {
	var c SimpleActions

	c.tag = tagSA.NewSimpleActions(
		others.entityIDGenerator,
		factories.tag,
		gateways.tag,
	)
	c.tagsSpace = tagsSpaceSA.NewSimpleActions(
		others.entityIDGenerator,
		factories.tagsSpace,
		gateways.tagsSpace, // Using gateway instead of AtomicAction because of the same interface
	)

	return &c
}
