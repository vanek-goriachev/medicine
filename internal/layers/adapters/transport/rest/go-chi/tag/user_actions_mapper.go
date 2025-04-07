//nolint:unparam // Required signatures for handlers generation
package tag

import (
	"fmt"
	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	dto "medicine/internal/layers/transport/rest/go-chi/tag"
	entityID "medicine/pkg/entity-id"
)

type UserActionsChiMapper struct {
	entityIDMapper entityID.Mapper
	tagChiMapper   tagChiMapper
}

func NewUserActionsChiMapper(
	entityIDMapper entityID.Mapper,
	tagChiMapper tagChiMapper,
) *UserActionsChiMapper {
	return &UserActionsChiMapper{
		entityIDMapper: entityIDMapper,
		tagChiMapper:   tagChiMapper,
	}
}

func (m *UserActionsChiMapper) TagForceCreateInFromChi(
	in dto.TagForceCreateIn,
) (tagUA.TagForceCreateIn, error) {
	tagsSpaceID, err := m.entityIDMapper.FromString(in.TagsSpaceID)
	if err != nil {
		return tagUA.TagForceCreateIn{}, fmt.Errorf("can't convert tags space id: %w", err)
	}

	return tagUA.TagForceCreateIn{
		Name:        in.Name,
		TagsSpaceID: tagsSpaceID,
	}, nil
}

func (m *UserActionsChiMapper) TagForceCreateOutToChi(out tagUA.TagForceCreateOut) dto.TagForceCreateOut {
	return dto.TagForceCreateOut{
		Tag: m.tagChiMapper.ToChi(out.Tag),
	}
}

//func (m *UserActionsChiMapper) TagGetByIDInFromChi(
//	in dto.TagGetByIDIn,
//) (tagUA.TagGetByIDIn, error) {
//	id, err := m.entityIDMapper.FromString(in.ID)
//	if err != nil {
//		return tagUA.TagGetByIDIn{}, fmt.Errorf("can't convert tags space id: %w", err)
//	}
//
//	return tagUA.TagGetByIDIn{
//		ID: id,
//	}, nil
//}
//
//func (m *UserActionsChiMapper) TagGetByIDOutToChi(out tagUA.TagGetByIDOut) dto.TagGetByIDOut {
//	return dto.TagGetByIDOut{
//		Tag: m.tagsSpaceChiMapper.ToChi(out.Tag),
//	}
//}
