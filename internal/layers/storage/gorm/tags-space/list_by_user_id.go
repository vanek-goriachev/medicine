package tags_space

import (
	"context"
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
)

func (g *GORMGateway) ListByUserID(
	_ context.Context,
	userID entityID.EntityID,
) ([]tagsSpaceModels.TagsSpace, error) {
	var tagsSpaces []TagsSpace

	result := g.db.Find(&tagsSpaces, "user_id = ?", userID)
	if result.Error != nil {
		return nil, fmt.Errorf("error filtering tagsSpaces by user_id: %w", result.Error)
	}

	return g.mapper.MultipleFromGORM(tagsSpaces), nil
}
