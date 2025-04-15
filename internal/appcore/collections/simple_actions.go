package collections

import (
	"medicine/internal/layers/business-logic/authorization"
	tagSA "medicine/internal/layers/business-logic/simple-actions/tag"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
)

type SimpleActions struct {
	tag       *tagSA.SimpleActions
	tagsSpace *tagsSpaceSA.SimpleActions
}

func NewSimpleActions(
	authorizer authorization.Authorizer,
	others *Others,
	gateways *DBGateways, // Using gateway instead of AtomicAction because of the same interface
	factories *Factories,
) *SimpleActions {
	var c SimpleActions

	c.tag = tagSA.NewSimpleActions(
		others.entityIDGenerator,
		factories.tag,
		gateways.tag,
	)
	c.tagsSpace = tagsSpaceSA.NewSimpleActions(
		authorizer,
		others.entityIDGenerator,
		factories.tagsSpace,
		gateways.tag,
		gateways.tagsSpace,
	)

	return &c
}
