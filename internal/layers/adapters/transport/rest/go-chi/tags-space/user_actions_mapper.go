package tags_space

import (
	createUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	dto "medicine/internal/layers/transport/rest/go-chi/tags-space"
)

type UserActionsChiMapper struct {
	tagsSpaceChiMapper tagsSpaceChiMapper
}

func NewUserActionsChiMapper(tagsSpaceChiMapper tagsSpaceChiMapper) *UserActionsChiMapper {
	return &UserActionsChiMapper{tagsSpaceChiMapper: tagsSpaceChiMapper}
}

func (*UserActionsChiMapper) CreateTagsSpaceInFromChi(
	in dto.CreateTagsSpaceIn,
) (createUA.CreateTagsSpaceIn, error) { //nolint:unparam // Required signature for handlers generation
	return createUA.CreateTagsSpaceIn{
		Name: in.Name,
	}, nil
}

func (m *UserActionsChiMapper) CreateTagsSpaceOutToChi(out createUA.CreateTagsSpaceOut) dto.CreateTagsSpaceOut {
	return dto.CreateTagsSpaceOut{
		TagsSpace: m.tagsSpaceChiMapper.ToChi(out.TagsSpace),
	}
}
