package medical_file

import (
	"context"
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type AtomicActions interface {
	Create(ctx context.Context, medicalFile medicalFileModels.MedicalFile) error
}

type MedicalFileFactory interface {
	New(
		id entityID.EntityID,
		name string,
		data medicalFileModels.DataType,
	) (medicalFileModels.MedicalFile, error)
}
