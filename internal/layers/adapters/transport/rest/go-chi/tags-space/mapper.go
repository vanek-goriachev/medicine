package tags_space

import (
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceChi "medicine/internal/layers/transport/rest/go-chi/tags-space"
	entityID "medicine/pkg/entity-id"
)

type ChiMapper struct {
	entityIDMapper entityID.Mapper
	tagChiMapper   tagChiMapper
}

func NewChiMapper(
	entityIDMapper entityID.Mapper,
	tagChiMapper tagChiMapper,
) *ChiMapper {
	return &ChiMapper{
		entityIDMapper: entityIDMapper,
		tagChiMapper:   tagChiMapper,
	}
}

func (m *ChiMapper) FromChi(chiTagsSpace tagsSpaceChi.TagsSpace) (tagsSpaceModels.TagsSpace, error) {
	id, err := m.entityIDMapper.FromString(chiTagsSpace.ID)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't convert tagsSpace id: %w", err)
	}

	tags, err := m.tagChiMapper.MultipleFromChi(chiTagsSpace.Tags)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't convert tags: %w", err)
	}

	return tagsSpaceModels.TagsSpace{
		ID:   id,
		Name: chiTagsSpace.Name,
		Tags: tags,
	}, nil
}

func (m *ChiMapper) MultipleFromChi(chiTagsSpaces []tagsSpaceChi.TagsSpace) ([]tagsSpaceModels.TagsSpace, error) {
	var err error
	tagsSpaces := make([]tagsSpaceModels.TagsSpace, len(chiTagsSpaces))

	for i, chiTagsSpace := range chiTagsSpaces {
		tagsSpaces[i], err = m.FromChi(chiTagsSpace)
		if err != nil {
			return nil, err
		}
	}

	return tagsSpaces, nil
}

func (m *ChiMapper) ToChi(tagsSpace tagsSpaceModels.TagsSpace) tagsSpaceChi.TagsSpace {
	return tagsSpaceChi.TagsSpace{
		ID:   tagsSpace.ID.String(),
		Name: tagsSpace.Name,
		Tags: m.tagChiMapper.MultipleToChi(tagsSpace.Tags),
	}
}

func (m *ChiMapper) MultipleToChi(tagsSpaces []tagsSpaceModels.TagsSpace) []tagsSpaceChi.TagsSpace {
	chiTagsSpaces := make([]tagsSpaceChi.TagsSpace, len(tagsSpaces))

	for i, tagsSpace := range tagsSpaces {
		chiTagsSpaces[i] = m.ToChi(tagsSpace)
	}

	return chiTagsSpaces
}
