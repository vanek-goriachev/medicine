package collections

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
)

type Factories struct {
	tagsSpace *tagsSpaceModels.Factory
}

func NewFactories() *Factories {
	var c Factories

	c.tagsSpace = tagsSpaceModels.NewFactory()

	return &c
}
