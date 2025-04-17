package collections

import (
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
)

type Factories struct {
	tag       *tagModels.Factory
	tagsSpace *tagsSpaceModels.Factory

	medicalFile *medicalFileModels.Factory

	visitRecord *visitRecordModels.Factory
}

func NewFactories(validators *Validators) *Factories {
	var c Factories

	c.tag = tagModels.NewFactory(validators.tag)
	c.tagsSpace = tagsSpaceModels.NewFactory(validators.tagsSpace)

	c.medicalFile = medicalFileModels.NewFactory(validators.medicalFile)

	c.visitRecord = visitRecordModels.NewFactory(validators.visitRecord)

	return &c
}
