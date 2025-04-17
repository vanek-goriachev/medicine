package visit_record

import (
	"context"
	"time"

	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
)

type SimpleActions interface {
	CreateWithEntities(
		ctx context.Context,
		name string,
		datetime time.Time,

		tagIDs []entityID.EntityID,
	) (visitRecordModels.VisitRecord, visitRecordModels.VisitRecordLinkedEntities, error)

	AttachMedicalFiles(
		ctx context.Context,
		visitRecordID entityID.EntityID,
		uploadedMedicalFiles []medicalFileModels.UploadedMedicalFile,
	) error
}
