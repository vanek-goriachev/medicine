package generators

import (
	tagsModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
)

func GenerateTagsSpace() tagsSpaceModels.TagsSpace {
	tagsSpaceID := GenerateEntityID()

	return tagsSpaceModels.TagsSpace{
		Name: "TagsSpace",
		ID:   tagsSpaceID,
		Tags: []tagsModels.Tag{
			GenerateTag(tagsSpaceID),
			GenerateTag(tagsSpaceID),
			GenerateTag(tagsSpaceID),
		},
	}
}
