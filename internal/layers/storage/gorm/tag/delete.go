package tag

import (
	"context"
	"fmt"
	gormModels "medicine/internal/layers/storage/gorm/models"
	entityID "medicine/pkg/entity-id"
)

func (g *GORMGateway) DeleteByID(_ context.Context, tagID entityID.EntityID) error {
	result := g.db.Model(gormModels.TagModel).Delete(nil, "id = ?", tagID)
	if result.Error != nil {
		return fmt.Errorf("error deleting tag by id: %w", result.Error)
	}

	return nil
}
