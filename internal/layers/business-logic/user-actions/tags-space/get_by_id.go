package tags_space

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type TagsSpaceGetByIDIn struct {
	ID entityID.EntityID
}

type TagsSpaceGetByIDOut struct {
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

func (ua *GetByIDUA) Act(
	ctx context.Context,
	user userModels.User,
	in TagsSpaceGetByIDIn,
) (TagsSpaceGetByIDOut, error) {
	tagsSpace, err := ua.simpleActions.GetByID(ctx, in.ID)
	if err != nil {
		return TagsSpaceGetByIDOut{}, fmt.Errorf("can't GetByID tags space: %w", err)
	}

	err = ua.authorizer.Authorize(
		ctx,
		user,
		authorization.GetTagsSpacePermission,
		authorization.TagsSpaceResource,
		tagsSpace.ID.String(),
	)
	if err != nil {
		return TagsSpaceGetByIDOut{}, authorization.NewUnauthorizedError(err)
	}

	return TagsSpaceGetByIDOut{
		TagsSpace: tagsSpace,
	}, nil
}
