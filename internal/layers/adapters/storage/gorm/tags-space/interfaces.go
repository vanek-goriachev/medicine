package tags_space

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagGORM "medicine/internal/layers/storage/gorm/tag"
)

type tagGORMMapper interface {
	FromGORM(dbTag tagGORM.Tag) tagModels.Tag
	MultipleFromGORM(dbTags []tagGORM.Tag) []tagModels.Tag
	ToGORM(tag tagModels.Tag) tagGORM.Tag
	MultipleToGORM(tags []tagModels.Tag) []tagGORM.Tag
}
