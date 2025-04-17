package visit_record

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
)

func (g *GORMGateway) LinkMedicalFiles(
	_ context.Context,
	visitRecordID entityID.EntityID,
	medicalFileIDs []entityID.EntityID,
) error {
	links := g.buildMedicalFilesLinks(visitRecordID, medicalFileIDs)

	result := g.db.Model(gormModels.VisitRecordMedicalFileModel).CreateInBatches(links, len(links))
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to insert links visitRecord-to-medicalFiles in DB: %w", err)
	}

	return nil
}

func (*GORMGateway) buildMedicalFilesLinks(
	visitRecordID entityID.EntityID,
	medicalFileIDs []entityID.EntityID,
) []gormModels.VisitRecordMedicalFile {
	links := make([]gormModels.VisitRecordMedicalFile, len(medicalFileIDs))
	for i, id := range medicalFileIDs {
		links[i] = gormModels.VisitRecordMedicalFile{
			VisitRecordID: uuid.UUID(visitRecordID),
			MedicalFileID: uuid.UUID(id),
		}
	}

	return links
}
