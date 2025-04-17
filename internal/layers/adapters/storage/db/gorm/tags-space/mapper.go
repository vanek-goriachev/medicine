package tags_space

import (
	"github.com/google/uuid"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
)

type GORMMapper struct {
	tagMapper tagGORMMapper
}

func NewGORMMapper(tagMapper tagGORMMapper) *GORMMapper {
	return &GORMMapper{tagMapper: tagMapper}
}

func (m *GORMMapper) FromGORM(dbTagsSpace gormModels.TagsSpace) tagsSpaceModels.TagsSpace {
	return tagsSpaceModels.TagsSpace{
		ID:   entityID.EntityID(dbTagsSpace.ID),
		Name: dbTagsSpace.Name,
		Tags: m.tagMapper.MultipleFromGORM(dbTagsSpace.Tags),
	}
}

func (m *GORMMapper) MultipleFromGORM(dbTagsSpaces []gormModels.TagsSpace) []tagsSpaceModels.TagsSpace {
	tagSpaces := make([]tagsSpaceModels.TagsSpace, len(dbTagsSpaces))
	for i, dbTagsSpace := range dbTagsSpaces {
		tagSpaces[i] = m.FromGORM(dbTagsSpace)
	}

	return tagSpaces
}

func (m *GORMMapper) ToGORM(tagsSpace tagsSpaceModels.TagsSpace) gormModels.TagsSpace {
	return gormModels.TagsSpace{
		ID:   uuid.UUID(tagsSpace.ID),
		Name: tagsSpace.Name,
		Tags: m.tagMapper.MultipleToGORM(tagsSpace.Tags),
	}
}

func (m *GORMMapper) MultipleToGORM(tagsSpaces []tagsSpaceModels.TagsSpace) []gormModels.TagsSpace {
	tagSpaces := make([]gormModels.TagsSpace, len(tagsSpaces))
	for i, tagsSpace := range tagsSpaces {
		tagSpaces[i] = m.ToGORM(tagsSpace)
	}

	return tagSpaces
}
