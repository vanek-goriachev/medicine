package collections

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
)

type Validators struct {
	entityID *entityID.Validator

	tag       *tagModels.Validator
	tagsSpace *tagsSpaceModels.Validator
}

func NewValidators() *Validators {
	var v Validators

	v.entityID = entityID.NewValidator()
	v.tag = tagModels.NewValidator(v.entityID)
	v.tagsSpace = tagsSpaceModels.NewValidator(v.entityID, v.tag)

	return &v
}
