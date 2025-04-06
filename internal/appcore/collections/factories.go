package collections

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
)

type Factories struct {
	tag       *tagModels.Factory
	tagsSpace *tagsSpaceModels.Factory
}

func NewFactories(validators *Validators) *Factories {
	var c Factories

	c.tag = tagModels.NewFactory(validators.tag)
	c.tagsSpace = tagsSpaceModels.NewFactory(validators.tagsSpace)

	return &c
}
