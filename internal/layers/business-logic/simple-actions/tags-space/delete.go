package tags_space

import (
	"context"
	"fmt"
	tags_space "medicine/internal/layers/business-logic/models/tags-space"

	entityID "medicine/pkg/entity-id"
)

func (sa *SimpleActions) Delete(
	ctx context.Context,
	id entityID.EntityID,
) error {
	_, err := sa.tagsSpaceAtomicActions.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("can't get tags_space: %w", err)
	}

	tags, err := sa.tagAtomicActions.FilterByTagsSpaceID(ctx, id)
	if err != nil {
		return fmt.Errorf("can't filter tags by tags_space_id: %w", err)
	} else if len(tags) > 0 {
		return tags_space.NewTagsSpaceHaveTagsError(id)
	}

	err = sa.tagsSpaceAtomicActions.DeleteByID(ctx, id)
	if err != nil {
		return fmt.Errorf("can't delete tags_space: %w", err)
	}

	return nil
}
