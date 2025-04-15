package tags_space

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	entityID "medicine/pkg/entity-id"
)

type TagsSpace struct {
	Name string
	Tags []tagModels.Tag
	ID   entityID.EntityID
}
