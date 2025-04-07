package generators

import (
	tagsModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
)

func GenerateTagsSpace(userID entityID.EntityID) tagsSpaceModels.TagsSpace {
	tagsSpaceID := GenerateEntityID()

	return tagsSpaceModels.TagsSpace{
		Name:   "TagsSpace",
		UserID: userID,
		ID:     tagsSpaceID,
		Tags: []tagsModels.Tag{
			GenerateTag(tagsSpaceID),
			GenerateTag(tagsSpaceID),
			GenerateTag(tagsSpaceID),
		},
	}
}
