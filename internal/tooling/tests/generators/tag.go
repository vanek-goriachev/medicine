package generators

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	entityID "medicine/pkg/entity-id"
)

func GenerateTag(tagsSpaceID entityID.EntityID) tagModels.Tag {
	return tagModels.Tag{
		Name:        RandomString(7), //nolint:mnd // Test data generation
		ID:          GenerateEntityID(),
		TagsSpaceID: tagsSpaceID,
	}
}
