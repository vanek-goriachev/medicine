package tags_space

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	gormModels "medicine/internal/layers/storage/gorm/models"
)

type tagsSpaceGORMMapper interface {
	FromGORM(dbTagsSpace gormModels.TagsSpace) tagsSpaceModels.TagsSpace
	MultipleFromGORM(dbTagsSpace []gormModels.TagsSpace) []tagsSpaceModels.TagsSpace
	ToGORM(tagsSpace tagsSpaceModels.TagsSpace) gormModels.TagsSpace
	MultipleToGORM(tagsSpace []tagsSpaceModels.TagsSpace) []gormModels.TagsSpace
}
