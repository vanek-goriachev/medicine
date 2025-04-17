//nolint:unparam // Required signatures for handlers generation
package visit_record

import (
	"fmt"
	visitRecordUA "medicine/internal/layers/business-logic/user-actions/visit-record"
	dto "medicine/internal/layers/transport/rest/go-chi/visit-record"
	datetimeMapper "medicine/internal/tooling/datetime"
	entityID "medicine/pkg/entity-id"
)

type UserActionsChiMapper struct {
	datetimeMapper       datetimeMapper.Mapper
	entityIDMapper       entityID.Mapper
	medicalFileChiMapper medicalFileChiMapper
	visitRecordChiMapper visitRecordChiMapper
}

func NewUserActionsChiMapper(
	datetimeMapper datetimeMapper.Mapper,
	entityIDMapper entityID.Mapper,
	medicalFileChiMapper medicalFileChiMapper,
	visitRecordChiMapper visitRecordChiMapper,
) *UserActionsChiMapper {
	return &UserActionsChiMapper{
		datetimeMapper:       datetimeMapper,
		entityIDMapper:       entityIDMapper,
		medicalFileChiMapper: medicalFileChiMapper,
		visitRecordChiMapper: visitRecordChiMapper,
	}
}

func (m *UserActionsChiMapper) VisitRecordCreateInFromChi(
	in dto.VisitRecordCreateIn,
) (visitRecordUA.VisitRecordCreateIn, error) {
	datetime, err := m.datetimeMapper.FromString(in.Datetime)
	if err != nil {
		return visitRecordUA.VisitRecordCreateIn{}, fmt.Errorf("cant parse datetime: %w", err)
	}

	tagIDs := make([]entityID.EntityID, len(in.TagIDs))
	for i, tagID := range in.TagIDs {
		tagIDs[i], err = m.entityIDMapper.FromString(tagID)
		if err != nil {
			return visitRecordUA.VisitRecordCreateIn{}, fmt.Errorf("cant map tagID: %w", err)
		}
	}

	return visitRecordUA.VisitRecordCreateIn{
		Name:     in.Name,
		Datetime: datetime,
		TagIDs:   tagIDs,
	}, nil
}

func (m *UserActionsChiMapper) VisitRecordCreateOutToChi(
	out visitRecordUA.VisitRecordCreateOut,
) dto.VisitRecordCreateOut {
	return dto.VisitRecordCreateOut{
		VisitRecord:               m.visitRecordChiMapper.ToChi(out.VisitRecord),
		VisitRecordLinkedEntities: m.visitRecordChiMapper.LinkedEntitiesToChi(out.VisitRecordLinkedEntities),
	}
}
