package tags_space

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type TagsSpaceListByUserIn struct{}

type TagsSpaceListByUserOut struct {
	TagsSpaces []tagsSpaceModels.TagsSpace
}

type ListByUserUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewListByUserUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *ListByUserUA {
	return &ListByUserUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *ListByUserUA) Act(
	ctx context.Context,
	user userModels.User,
	_ TagsSpaceListByUserIn,
) (TagsSpaceListByUserOut, error) {
	tagsSpaces, err := ua.simpleActions.ListByUser(ctx, user)
	if err != nil {
		return TagsSpaceListByUserOut{}, fmt.Errorf("can't list tags spaces by user (ua): %w", err)
	}

	err = ua.authorize(ctx, user, tagsSpaces)
	if err != nil {
		return TagsSpaceListByUserOut{}, err
	}

	return TagsSpaceListByUserOut{
		TagsSpaces: tagsSpaces,
	}, nil
}

func (ua *ListByUserUA) authorize(
	ctx context.Context,
	user userModels.User,
	tagsSpaces []tagsSpaceModels.TagsSpace,
) error {
	actions := make([]authorization.Action, len(tagsSpaces))
	for i, tagsSpace := range tagsSpaces {
		actions[i] = authorization.NewAction(
			authorization.ReadTagsSpacePermission,
			authorization.TagsSpaceResource,
			tagsSpace.ID.String(),
		)
	}

	err := ua.authorizer.BatchAuthorize(
		ctx,
		user,
		actions,
	)
	if err != nil {
		return fmt.Errorf("unathorized to list tags spaces for user (ua): %w", err)
	}

	return nil
}
