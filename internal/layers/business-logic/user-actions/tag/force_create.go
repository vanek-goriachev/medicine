package tag

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	tagModels "medicine/internal/layers/business-logic/models/tag"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type TagForceCreateIn struct {
	Name        string
	TagsSpaceID entityID.EntityID
}

type TagForceCreateOut struct {
	Tag tagModels.Tag
}

type ForceCreateUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewForceCreateUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *ForceCreateUA {
	return &ForceCreateUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *ForceCreateUA) Act(
	ctx context.Context,
	user userModels.User,
	in *TagForceCreateIn,
) (TagForceCreateOut, error) {
	err := ua.authorizer.Authorize(
		ctx,
		user,
		authorization.NewAction(
			authorization.CreateTagPermission,
			authorization.TagResource,
			entityID.EntityID{},
		),
	)
	if err != nil {
		return TagForceCreateOut{}, authorization.NewUnauthorizedError(err)
	}

	tag, err := ua.simpleActions.Create(ctx, in.Name, in.TagsSpaceID)
	if err != nil {
		return TagForceCreateOut{}, fmt.Errorf("can't create tag (ua): %w", err)
	}

	return TagForceCreateOut{
		Tag: tag,
	}, nil
}
