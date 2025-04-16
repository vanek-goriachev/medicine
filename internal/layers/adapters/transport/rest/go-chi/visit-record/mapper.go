package visit_record

import (
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	visitRecordChi "medicine/internal/layers/transport/rest/go-chi/visit-record"
	"medicine/internal/tooling/datetime"
	entityID "medicine/pkg/entity-id"
)

type ChiMapper struct {
	datetimeMapper datetime.Mapper
	entityIDMapper entityID.Mapper
}

func NewChiMapper(
	datetimeMapper datetime.Mapper,
	entityIDMapper entityID.Mapper,
) *ChiMapper {
	return &ChiMapper{
		datetimeMapper: datetimeMapper,
		entityIDMapper: entityIDMapper,
	}
}

func (m *ChiMapper) ToChi(visitRecord visitRecordModels.VisitRecord) visitRecordChi.VisitRecord {
	return visitRecordChi.VisitRecord{
		ID:       visitRecord.ID.String(),
		Name:     visitRecord.Name,
		Datetime: m.datetimeMapper.ToString(visitRecord.Datetime),
	}
}

func (m *ChiMapper) LinkedEntitiesToChi(
	visitRecordLinkedEntities visitRecordModels.VisitRecordLinkedEntities,
) visitRecordChi.VisitRecordLinkedEntities {
	tagIDs := make([]string, len(visitRecordLinkedEntities.TagIDs))
	for i, tagID := range visitRecordLinkedEntities.TagIDs {
		tagIDs[i] = tagID.String()
	}

	medicalFileIDs := make([]string, len(visitRecordLinkedEntities.MedicalFileIDs))
	for i, medicalFileID := range visitRecordLinkedEntities.MedicalFileIDs {
		medicalFileIDs[i] = medicalFileID.String()
	}

	return visitRecordChi.VisitRecordLinkedEntities{
		TagIDs:         tagIDs,
		MedicalFileIDs: medicalFileIDs,
	}
}
