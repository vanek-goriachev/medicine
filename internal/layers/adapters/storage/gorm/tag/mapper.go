package tag

import (
	"github.com/google/uuid"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagGORM "medicine/internal/layers/storage/gorm/tag"
	entityID "medicine/pkg/entity-id"
)

type GORMMapper struct{}

func NewGORMMapper() *GORMMapper {
	return &GORMMapper{}
}

func (*GORMMapper) FromGORM(dbTag tagGORM.Tag) tagModels.Tag {
	return tagModels.Tag{
		ID:          entityID.EntityID(dbTag.ID),
		TagsSpaceID: entityID.EntityID(dbTag.TagsSpaceID),
		Name:        dbTag.Name,
	}
}

func (t *GORMMapper) MultipleFromGORM(dbTags []tagGORM.Tag) []tagModels.Tag {
	tags := make([]tagModels.Tag, len(dbTags))
	for i, dbTag := range dbTags {
		tags[i] = t.FromGORM(dbTag)
	}

	return tags
}

func (*GORMMapper) ToGORM(tag tagModels.Tag) tagGORM.Tag {
	return tagGORM.Tag{
		ID:          uuid.UUID(tag.ID),
		TagsSpaceID: uuid.UUID(tag.TagsSpaceID),
		Name:        tag.Name,
	}
}

func (t *GORMMapper) MultipleToGORM(tags []tagModels.Tag) []tagGORM.Tag {
	dbTags := make([]tagGORM.Tag, len(tags))
	for i, tag := range tags {
		dbTags[i] = t.ToGORM(tag)
	}

	return dbTags
}
