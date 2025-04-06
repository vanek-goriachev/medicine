package tags_space

import (
	"context"
	"fmt"
	entityID "medicine/pkg/entity-id"

	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type GetByIDTagsSpaceIn struct {
	ID entityID.EntityID
}

type GetByIDTagsSpaceOut struct {
	TagsSpace tagsSpaceModels.TagsSpace
}

type GetByIDUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewGetByIDUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *GetByIDUA {
	return &GetByIDUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *GetByIDUA) Act(ctx context.Context, user userModels.User, in GetByIDTagsSpaceIn) (GetByIDTagsSpaceOut, error) {
	tagsSpace, err := ua.simpleActions.GetByID(ctx, in.ID)
	if err != nil {
		return GetByIDTagsSpaceOut{}, fmt.Errorf("can't GetByID tags space: %w", err)
	}

	err = ua.authorizer.Authorize(
		ctx,
		user,
		authorization.GetTagsSpacePermission,
		authorization.TagsSpaceResource,
		tagsSpace.ID.String(),
	)
	if err != nil {
		return GetByIDTagsSpaceOut{}, fmt.Errorf("authorization failed: %w", err)
	}

	return GetByIDTagsSpaceOut{
		TagsSpace: tagsSpace,
	}, nil
}
