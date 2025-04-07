package tag

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	gormModels "medicine/internal/layers/storage/gorm/models"
)

type tagGORMMapper interface {
	FromGORM(dbTag gormModels.Tag) tagModels.Tag
	MultipleFromGORM(dbTags []gormModels.Tag) []tagModels.Tag
	ToGORM(tag tagModels.Tag) gormModels.Tag
	MultipleToGORM(tags []tagModels.Tag) []gormModels.Tag
}
