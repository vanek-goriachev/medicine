package visit_record

import (
	"context"
	"time"

	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type AtomicActions interface {
	Create(ctx context.Context, visitRecord visitRecordModels.VisitRecord) error
	LinkTags(ctx context.Context, visitRecord entityID.EntityID, tagIDs []entityID.EntityID) error
	LinkMedicalFiles(ctx context.Context, visitRecordID entityID.EntityID, medicalFileIDs []entityID.EntityID) error
}

type MedicalFileAtomicActions interface {
	CreateMultiple(
		ctx context.Context,
		uploadedFiles []medicalFileModels.UploadedMedicalFile,
	) ([]medicalFileModels.MedicalFile, error)
}

type VisitRecordFactory interface {
	New(
		id entityID.EntityID,
		name string,
		datetime time.Time,
	) (visitRecordModels.VisitRecord, error)
}
