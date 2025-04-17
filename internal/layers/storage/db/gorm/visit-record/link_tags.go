package visit_record

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
)

func (g *GORMGateway) LinkTags(
	_ context.Context,
	visitRecordID entityID.EntityID,
	tagIDs []entityID.EntityID,
) error {
	links := g.buildTagsLinks(visitRecordID, tagIDs)

	result := g.db.Model(&gormModels.VisitRecordTag{}).CreateInBatches(links, len(links))
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to insert links visitRecord-to-tags in DB: %w", err)
	}

	return nil
}

func (g *GORMGateway) buildTagsLinks(
	visitRecordID entityID.EntityID,
	tagIDs []entityID.EntityID,
) []gormModels.VisitRecordTag {
	links := make([]gormModels.VisitRecordTag, len(tagIDs))
	for i, id := range tagIDs {
		links[i] = gormModels.VisitRecordTag{
			VisitRecordID: uuid.UUID(visitRecordID),
			TagID:         uuid.UUID(id),
		}
	}

	return links
}
