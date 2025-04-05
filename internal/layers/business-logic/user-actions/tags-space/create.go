package tags_space

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type CreateTagsSpaceIn struct {
	Name string
}

type CreateTagsSpaceOut struct {
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

func (ua *CreateUA) Act(ctx context.Context, in CreateTagsSpaceIn) (CreateTagsSpaceOut, error) {
	user, err := userModels.GetFromContext(ctx)
	if err != nil {
		return CreateTagsSpaceOut{}, fmt.Errorf("can't get user from context: %w", err)
	}

	err = ua.authorizer.Authorize(
		ctx,
		user,
		authorization.CreateTagsSpacePermission,
		authorization.TagsSpaceResource,
		"",
	)
	if err != nil {
		return CreateTagsSpaceOut{}, fmt.Errorf("authorization failed: %w", err)
	}

	tagsSpace, err := ua.simpleActions.Create(ctx, user, in.Name)
	if err != nil {
		return CreateTagsSpaceOut{}, fmt.Errorf("can't create tags space: %w", err)
	}

	return CreateTagsSpaceOut{
		TagsSpace: tagsSpace,
	}, nil
}
