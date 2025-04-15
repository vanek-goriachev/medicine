package tags_space

import (
	"context"
	"fmt"

	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
)

func (g *GORMGateway) DeleteByID(_ context.Context, tagsSpaceID entityID.EntityID) error {
	result := g.db.Model(gormModels.TagsSpaceModel).Delete(nil, WhereTagsSpaceIDEquals, tagsSpaceID)
	if result.Error != nil {
		return fmt.Errorf("error deleting tags_space by id: %w", result.Error)
	}

	return nil
}
