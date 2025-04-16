package collections

import (
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
)

type Validators struct {
	entityID *entityID.Validator

	tag       *tagModels.Validator
	tagsSpace *tagsSpaceModels.Validator

	medicalFile *medicalFileModels.Validator

	visitRecord *visitRecordModels.Validator
}

func NewValidators() *Validators {
	var v Validators

	v.entityID = entityID.NewValidator()

	v.tag = tagModels.NewValidator(v.entityID)
	v.tagsSpace = tagsSpaceModels.NewValidator(v.entityID, v.tag)

	v.medicalFile = medicalFileModels.NewValidator(v.entityID)

	v.visitRecord = visitRecordModels.NewValidator(v.entityID)

	return &v
}
