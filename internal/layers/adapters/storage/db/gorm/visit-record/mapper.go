package visit_record

import (
	"github.com/google/uuid"

	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
)

type GORMMapper struct{}

func NewGORMMapper() *GORMMapper {
	return &GORMMapper{}
}

func (*GORMMapper) FromGORM(dbVisitRecord gormModels.VisitRecord) visitRecordModels.VisitRecord {
	return visitRecordModels.VisitRecord{
		ID:       entityID.EntityID(dbVisitRecord.ID),
		Name:     dbVisitRecord.Name,
		Datetime: dbVisitRecord.Datetime,
	}
}

func (*GORMMapper) ToGORM(visitRecord visitRecordModels.VisitRecord) gormModels.VisitRecord {
	return gormModels.VisitRecord{
		ID:       uuid.UUID(visitRecord.ID),
		Name:     visitRecord.Name,
		Datetime: visitRecord.Datetime,
	}
}
