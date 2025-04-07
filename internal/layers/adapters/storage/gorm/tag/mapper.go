package tag

import (
	"github.com/google/uuid"
	gormModels "medicine/internal/layers/storage/gorm/models"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	entityID "medicine/pkg/entity-id"
)

type GORMMapper struct{}

func NewGORMMapper() *GORMMapper {
	return &GORMMapper{}
}

func (*GORMMapper) FromGORM(dbTag gormModels.Tag) tagModels.Tag {
	return tagModels.Tag{
		ID:          entityID.EntityID(dbTag.ID),
		TagsSpaceID: entityID.EntityID(dbTag.TagsSpaceID),
		Name:        dbTag.Name,
	}
}

func (t *GORMMapper) MultipleFromGORM(dbTags []gormModels.Tag) []tagModels.Tag {
	tags := make([]tagModels.Tag, len(dbTags))
	for i, dbTag := range dbTags {
		tags[i] = t.FromGORM(dbTag)
	}

	return tags
}

func (*GORMMapper) ToGORM(tag tagModels.Tag) gormModels.Tag {
	return gormModels.Tag{
		ID:          uuid.UUID(tag.ID),
		TagsSpaceID: uuid.UUID(tag.TagsSpaceID),
		Name:        tag.Name,
	}
}

func (t *GORMMapper) MultipleToGORM(tags []tagModels.Tag) []gormModels.Tag {
	dbTags := make([]gormModels.Tag, len(tags))
	for i, tag := range tags {
		dbTags[i] = t.ToGORM(tag)
	}

	return dbTags
}
