package tags_space

import (
	"context"
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
)

func (sa *SimpleActions) GetByID(
	ctx context.Context,
	id entityID.EntityID,
) (tagsSpaceModels.TagsSpace, error) {
	tagsSpace, err := sa.atomicActions.GetByID(ctx, id)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't get tagsSpace by ID: %w", err)
	}

	return tagsSpace, nil
}
