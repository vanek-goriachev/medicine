package tag

import (
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagChi "medicine/internal/layers/transport/rest/go-chi/tag"
	entityID "medicine/pkg/entity-id"
)

type ChiMapper struct {
	entityIDMapper entityID.Mapper
}

func NewChiMapper(entityIDMapper entityID.Mapper) *ChiMapper {
	return &ChiMapper{entityIDMapper: entityIDMapper}
}

func (m *ChiMapper) FromChi(chiTag tagChi.Tag) (tagModels.Tag, error) {
	id, err := m.entityIDMapper.FromString(chiTag.ID)
	if err != nil {
		return tagModels.Tag{}, fmt.Errorf("failed to convert tag id: %w", err)
	}

	tagsSpaceID, err := m.entityIDMapper.FromString(chiTag.TagsSpaceID)
	if err != nil {
		return tagModels.Tag{}, fmt.Errorf("failed to convert tags space id: %w", err)
	}

	return tagModels.Tag{
		ID:          id,
		Name:        chiTag.Name,
		TagsSpaceID: tagsSpaceID,
	}, nil
}

func (m *ChiMapper) MultipleFromChi(chiTags []tagChi.Tag) ([]tagModels.Tag, error) {
	var err error
	tags := make([]tagModels.Tag, len(chiTags))

	for i, chiTag := range chiTags {
		tags[i], err = m.FromChi(chiTag)
		if err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func (*ChiMapper) ToChi(tag tagModels.Tag) tagChi.Tag {
	return tagChi.Tag{
		ID:          tag.ID.String(),
		Name:        tag.Name,
		TagsSpaceID: tag.TagsSpaceID.String(),
	}
}

func (m *ChiMapper) MultipleToChi(tags []tagModels.Tag) []tagChi.Tag {
	chiTags := make([]tagChi.Tag, len(tags))

	for i, tag := range tags {
		chiTags[i] = m.ToChi(tag)
	}

	return chiTags
}
