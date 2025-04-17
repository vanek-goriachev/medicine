package visit_record

import (
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
)

type visitRecordGORMMapper interface {
	FromGORM(dbVisitRecord gormModels.VisitRecord) visitRecordModels.VisitRecord
	ToGORM(visitRecord visitRecordModels.VisitRecord) gormModels.VisitRecord
}
