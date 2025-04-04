package tag

import tagModels "medicine/internal/layers/business-logic/models/tag"

type tagGORMMapper interface {
	FromGORM(dbTag Tag) tagModels.Tag
	MultipleFromGORM(dbTags []Tag) []tagModels.Tag
	ToGORM(tag tagModels.Tag) Tag
	MultipleToGORM(tags []tagModels.Tag) []Tag
}
