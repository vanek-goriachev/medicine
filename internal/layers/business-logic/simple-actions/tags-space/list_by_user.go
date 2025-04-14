package tags_space

import (
	"context"
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

func (sa *SimpleActions) ListByUser(
	ctx context.Context,
	user userModels.User,
) ([]tagsSpaceModels.TagsSpace, error) {
	tagsSpace, err := sa.tagsSpaceAtomicActions.ListByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("can't get list tags spaces by user: %w", err)
	}

	return tagsSpace, nil
}
