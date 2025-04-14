package tag

import (
	"context"
	"fmt"

	entityID "medicine/pkg/entity-id"
)

func (sa *SimpleActions) UntagAllAndDelete(
	ctx context.Context,
	id entityID.EntityID,
) error {
	_, err := sa.atomicActions.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("can't get tag before deletion: %w", err)
	}

	err = sa.UntagAll(ctx, id)
	if err != nil {
		return fmt.Errorf("error when untagging resources: %w", err)
	}

	err = sa.atomicActions.DeleteByID(ctx, id)
	if err != nil {
		return fmt.Errorf("can't delete tag: %w", err)
	}

	return nil
}
