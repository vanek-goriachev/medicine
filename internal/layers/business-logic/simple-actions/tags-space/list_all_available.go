package tags_space

import (
	"context"
	"fmt"
	"medicine/internal/layers/business-logic/authorization"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

func (sa *SimpleActions) ListAllAvailable(
	ctx context.Context,
	user userModels.User,
) ([]tagsSpaceModels.TagsSpace, error) {
	ids, err := sa.authorizer.AvailableResources(
		ctx,
		user,
		authorization.TagsSpaceResource,
		authorization.ReadTagsSpacePermission,
	)
	switch {
	case err != nil:
		return nil, fmt.Errorf("can't get available tags spaces ids: %w", err)
	case len(ids) == 0:
		return []tagsSpaceModels.TagsSpace{}, nil
	}

	tagsSpace, err := sa.tagsSpaceAtomicActions.ListByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("can't get list all available tags spaces: %w", err)
	}

	return tagsSpace, nil
}
