package collections

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
)

type Factories struct {
	tagsSpace *tagsSpaceModels.Factory
}

func NewFactories(validators *Validators) *Factories {
	var c Factories

	c.tagsSpace = tagsSpaceModels.NewFactory(validators.tagsSpace)

	return &c
}
