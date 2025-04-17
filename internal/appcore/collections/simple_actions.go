package collections

import (
	"medicine/internal/layers/business-logic/authorization"
	medicalFileSA "medicine/internal/layers/business-logic/simple-actions/medical-file"
	tagSA "medicine/internal/layers/business-logic/simple-actions/tag"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	visitRecordSA "medicine/internal/layers/business-logic/simple-actions/visit-record"
)

type SimpleActions struct {
	tag       *tagSA.SimpleActions
	tagsSpace *tagsSpaceSA.SimpleActions

	medicalFile *medicalFileSA.SimpleActions

	visitRecord *visitRecordSA.SimpleActions
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

	c.medicalFile = medicalFileSA.NewSimpleActions(
		others.entityIDGenerator,
		factories.medicalFile,
		gateways.medicalFile,
	)

	c.visitRecord = visitRecordSA.NewSimpleActions(
		others.entityIDGenerator,
		factories.visitRecord,
		c.medicalFile,
		gateways.visitRecord,
	)

	return &c
}
