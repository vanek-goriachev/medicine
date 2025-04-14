//nolint:unparam // Required signatures for handlers generation
package tags_space

import (
	"fmt"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	dto "medicine/internal/layers/transport/rest/go-chi/tags-space"
	entityID "medicine/pkg/entity-id"
)

type UserActionsChiMapper struct {
	entityIDMapper     entityID.Mapper
	tagsSpaceChiMapper tagsSpaceChiMapper
}

func NewUserActionsChiMapper(
	entityIDMapper entityID.Mapper,
	tagsSpaceChiMapper tagsSpaceChiMapper,
) *UserActionsChiMapper {
	return &UserActionsChiMapper{
		entityIDMapper:     entityIDMapper,
		tagsSpaceChiMapper: tagsSpaceChiMapper,
	}
}

func (*UserActionsChiMapper) TagsSpaceCreateInFromChi(
	in dto.TagsSpaceCreateIn,
) (tagsSpaceUA.TagsSpaceCreateIn, error) {
	return tagsSpaceUA.TagsSpaceCreateIn{
		Name: in.Name,
	}, nil
}

func (m *UserActionsChiMapper) TagsSpaceCreateOutToChi(out tagsSpaceUA.TagsSpaceCreateOut) dto.TagsSpaceCreateOut {
	return dto.TagsSpaceCreateOut{
		TagsSpace: m.tagsSpaceChiMapper.ToChi(out.TagsSpace),
	}
}

func (*UserActionsChiMapper) TagsSpaceListByUserInFromChi(
	_ dto.TagsSpaceListByUserIn,
) (tagsSpaceUA.TagsSpaceListByUserIn, error) {
	return tagsSpaceUA.TagsSpaceListByUserIn{}, nil
}

func (m *UserActionsChiMapper) TagsSpaceListByUserOutToChi(
	out tagsSpaceUA.TagsSpaceListByUserOut,
) dto.TagsSpaceListByUserOut {
	return dto.TagsSpaceListByUserOut{
		TagsSpaces: m.tagsSpaceChiMapper.MultipleToChi(out.TagsSpaces),
	}
}

func (m *UserActionsChiMapper) TagsSpaceGetByIDInFromChi(
	in dto.TagsSpaceGetByIDIn,
) (tagsSpaceUA.TagsSpaceGetByIDIn, error) {
	id, err := m.entityIDMapper.FromString(in.ID)
	if err != nil {
		return tagsSpaceUA.TagsSpaceGetByIDIn{}, fmt.Errorf("can't convert tags space id: %w", err)
	}

	return tagsSpaceUA.TagsSpaceGetByIDIn{
		ID: id,
	}, nil
}

func (m *UserActionsChiMapper) TagsSpaceGetByIDOutToChi(out tagsSpaceUA.TagsSpaceGetByIDOut) dto.TagsSpaceGetByIDOut {
	return dto.TagsSpaceGetByIDOut{
		TagsSpace: m.tagsSpaceChiMapper.ToChi(out.TagsSpace),
	}
}

func (m *UserActionsChiMapper) TagsSpaceDeleteInFromChi(
	in dto.TagsSpaceDeleteIn,
) (tagsSpaceUA.TagsSpaceDeleteIn, error) {
	id, err := m.entityIDMapper.FromString(in.ID)
	if err != nil {
		return tagsSpaceUA.TagsSpaceDeleteIn{}, fmt.Errorf("can't convert tags space id: %w", err)
	}

	return tagsSpaceUA.TagsSpaceDeleteIn{
		ID: id,
	}, nil
}

func (m *UserActionsChiMapper) TagsSpaceDeleteOutToChi(_ tagsSpaceUA.TagsSpaceDeleteOut) dto.TagsSpaceDeleteOut {
	return dto.TagsSpaceDeleteOut{}
}
