package visit_record

import (
	"context"
	"fmt"
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	entityID "medicine/pkg/entity-id"
)

func (sa *SimpleActions) AttachMedicalFiles(
	ctx context.Context,
	visitRecordID entityID.EntityID,
	uploadedMedicalFiles []medicalFileModels.UploadedMedicalFile,
) error {
	_, err := sa.atomicActions.GetByID(ctx, visitRecordID)
	if err != nil {
		return fmt.Errorf("failed to get visit record: %w", err)
	}

	medicalFiles, err := sa.medicalFileAtomicActions.CreateMultiple(ctx, uploadedMedicalFiles)
	if err != nil {
		return fmt.Errorf("failed to create medical files: %w", err)
	}

	err = sa.atomicActions.LinkMedicalFiles(
		ctx,
		visitRecordID,
		sa.extractMedicalFilesIDs(medicalFiles),
	)
	if err != nil {
		return fmt.Errorf("failed to link medical files: %w", err)
	}

	return nil
}

func (sa *SimpleActions) extractMedicalFilesIDs(
	medicalFiles []medicalFileModels.MedicalFile,
) []entityID.EntityID {
	medicalFilesIDs := make([]entityID.EntityID, len(medicalFiles))
	for i, medicalFile := range medicalFiles {
		medicalFilesIDs[i] = medicalFile.ID
	}
	return medicalFilesIDs
}
