package tags_space

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	entityID "medicine/pkg/entity-id"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (*Factory) New(
	id entityID.EntityID,
	userID entityID.EntityID,
	name string,
	tags []tagModels.Tag,
) TagsSpace {
	return TagsSpace{
		ID:     id,
		UserID: userID,
		Name:   name,
		Tags:   tags,
	}
}
