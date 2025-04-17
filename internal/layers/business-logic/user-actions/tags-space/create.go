package tags_space

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type TagsSpaceCreateIn struct {
	Name string
}

type TagsSpaceCreateOut struct {
	TagsSpace tagsSpaceModels.TagsSpace
}

type CreateUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewCreateUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *CreateUA {
	return &CreateUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *CreateUA) Act(ctx context.Context, user userModels.User, in *TagsSpaceCreateIn) (TagsSpaceCreateOut, error) {
	err := ua.authorizer.Authorize(
		ctx,
		user,
		authorization.NewAction(
			authorization.CreateTagsSpacePermission,
			authorization.TagsSpaceResource,
			entityID.EntityID{},
		),
	)
	if err != nil {
		return TagsSpaceCreateOut{}, authorization.NewUnauthorizedError(err)
	}

	tagsSpace, err := ua.simpleActions.Create(ctx, user, in.Name)
	if err != nil {
		return TagsSpaceCreateOut{}, fmt.Errorf("can't create tags space (ua): %w", err)
	}

	return TagsSpaceCreateOut{
		TagsSpace: tagsSpace,
	}, nil
}
