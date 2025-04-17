package tags_space

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type TagsSpaceDeleteIn struct {
	ID entityID.EntityID
}

type TagsSpaceDeleteOut struct{}

type DeleteUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewDeleteUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *DeleteUA {
	return &DeleteUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *DeleteUA) Act(
	ctx context.Context,
	user userModels.User,
	in *TagsSpaceDeleteIn,
) (TagsSpaceDeleteOut, error) { //nolint:unparam // UserAction signature requires returned parameter
	err := ua.authorizer.Authorize(
		ctx,
		user,
		authorization.NewAction(
			authorization.DeleteTagsSpacePermission,
			authorization.TagsSpaceResource,
			in.ID,
		),
	)
	if err != nil {
		return TagsSpaceDeleteOut{}, authorization.NewUnauthorizedError(err)
	}

	err = ua.simpleActions.Delete(ctx, in.ID)
	if err != nil {
		return TagsSpaceDeleteOut{}, fmt.Errorf("can't delete tagsSpace (ua): %w", err)
	}

	return TagsSpaceDeleteOut{}, nil
}
