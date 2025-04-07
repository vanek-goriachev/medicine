package tag

import (
	"context"
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	gormModels "medicine/internal/layers/storage/gorm/models"
	entityID "medicine/pkg/entity-id"
)

func (g *GORMGateway) FilterByTagsSpaceID(_ context.Context, tagsSpaceID entityID.EntityID) ([]tagModels.Tag, error) {
	var dbTags []gormModels.Tag

	result := g.db.Find(&dbTags, "tags_space_id = ?", tagsSpaceID)
	if result.Error != nil {
		return nil, fmt.Errorf("error filtering tags by tags_space_id: %w", result.Error)
	}

	return g.mapper.MultipleFromGORM(dbTags), nil
}
