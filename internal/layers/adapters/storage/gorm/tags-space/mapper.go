package tags_space

import (
	"github.com/google/uuid"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceGORM "medicine/internal/layers/storage/gorm/tags-space"
	entityID "medicine/pkg/entity-id"
)

type GORMMapper struct {
	tagMapper tagGORMMapper
}

func NewGORMMapper(tagMapper tagGORMMapper) *GORMMapper {
	return &GORMMapper{tagMapper: tagMapper}
}

func (m *GORMMapper) FromGORM(dbTagsSpace tagsSpaceGORM.TagsSpace) tagsSpaceModels.TagsSpace {
	return tagsSpaceModels.TagsSpace{
		ID:     entityID.EntityID(dbTagsSpace.ID),
		UserID: entityID.EntityID(dbTagsSpace.UserID),
		Name:   dbTagsSpace.Name,
		Tags:   m.tagMapper.MultipleFromGORM(dbTagsSpace.Tags),
	}
}

func (m *GORMMapper) MultipleFromGORM(dbTagsSpaces []tagsSpaceGORM.TagsSpace) []tagsSpaceModels.TagsSpace {
	tagSpaces := make([]tagsSpaceModels.TagsSpace, len(dbTagsSpaces))
	for i, dbTagsSpace := range dbTagsSpaces {
		tagSpaces[i] = m.FromGORM(dbTagsSpace)
	}

	return tagSpaces
}

func (m *GORMMapper) ToGORM(tagsSpace tagsSpaceModels.TagsSpace) tagsSpaceGORM.TagsSpace {
	return tagsSpaceGORM.TagsSpace{
		ID:     uuid.UUID(tagsSpace.ID),
		UserID: uuid.UUID(tagsSpace.UserID),
		Name:   tagsSpace.Name,
		Tags:   m.tagMapper.MultipleToGORM(tagsSpace.Tags),
	}
}

func (m *GORMMapper) MultipleToGORM(tagsSpaces []tagsSpaceModels.TagsSpace) []tagsSpaceGORM.TagsSpace {
	tagSpaces := make([]tagsSpaceGORM.TagsSpace, len(tagsSpaces))
	for i, tagsSpace := range tagsSpaces {
		tagSpaces[i] = m.ToGORM(tagsSpace)
	}

	return tagSpaces
}
