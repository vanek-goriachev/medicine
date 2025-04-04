package tags_space

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
)

type tagsSpaceGORMMapper interface {
	FromGORM(dbTagsSpace TagsSpace) tagsSpaceModels.TagsSpace
	MultipleFromGORM(dbTagsSpace []TagsSpace) []tagsSpaceModels.TagsSpace
	ToGORM(tagsSpace tagsSpaceModels.TagsSpace) TagsSpace
	MultipleToGORM(tagsSpace []tagsSpaceModels.TagsSpace) []TagsSpace
}
