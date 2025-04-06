//nolint:unparam // Required signatures for handlers generation
package tags_space

import (
	"fmt"

	createUA "medicine/internal/layers/business-logic/user-actions/tags-space"
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
) (createUA.TagsSpaceCreateIn, error) {
	return createUA.TagsSpaceCreateIn{
		Name: in.Name,
	}, nil
}

func (m *UserActionsChiMapper) TagsSpaceCreateOutToChi(out createUA.TagsSpaceCreateOut) dto.TagsSpaceCreateOut {
	return dto.TagsSpaceCreateOut{
		TagsSpace: m.tagsSpaceChiMapper.ToChi(out.TagsSpace),
	}
}

func (*UserActionsChiMapper) TagsSpaceListByUserInFromChi(
	_ dto.TagsSpaceListByUserIn,
) (createUA.TagsSpaceListByUserIn, error) {
	return createUA.TagsSpaceListByUserIn{}, nil
}

func (m *UserActionsChiMapper) TagsSpaceListByUserOutToChi(
	out createUA.TagsSpaceListByUserOut,
) dto.TagsSpaceListByUserOut {
	return dto.TagsSpaceListByUserOut{
		TagsSpaces: m.tagsSpaceChiMapper.MultipleToChi(out.TagsSpaces),
	}
}

func (m *UserActionsChiMapper) TagsSpaceGetByIDInFromChi(
	in dto.TagsSpaceGetByIDIn,
) (createUA.TagsSpaceGetByIDIn, error) {
	id, err := m.entityIDMapper.FromString(in.ID)
	if err != nil {
		return createUA.TagsSpaceGetByIDIn{}, fmt.Errorf("can't convert tags space id: %w", err)
	}

	return createUA.TagsSpaceGetByIDIn{
		ID: id,
	}, nil
}

func (m *UserActionsChiMapper) TagsSpaceGetByIDOutToChi(out createUA.TagsSpaceGetByIDOut) dto.TagsSpaceGetByIDOut {
	return dto.TagsSpaceGetByIDOut{
		TagsSpace: m.tagsSpaceChiMapper.ToChi(out.TagsSpace),
	}
}
