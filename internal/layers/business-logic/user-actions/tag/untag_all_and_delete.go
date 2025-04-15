package tag

import (
	"context"
	"fmt"

	"medicine/internal/layers/business-logic/authorization"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type TagUntagAllAndDeleteIn struct {
	ID entityID.EntityID
}

type TagUntagAllAndDeleteOut struct{}

type UntagAllAndDeleteUA struct {
	authorizer    authorization.Authorizer
	simpleActions SimpleActions
}

func NewUntagAllAndDeleteUA(
	authorizer authorization.Authorizer,
	simpleActions SimpleActions,
) *UntagAllAndDeleteUA {
	return &UntagAllAndDeleteUA{
		simpleActions: simpleActions,
		authorizer:    authorizer,
	}
}

func (ua *UntagAllAndDeleteUA) Act(
	ctx context.Context,
	user userModels.User,
	in TagUntagAllAndDeleteIn,
) (TagUntagAllAndDeleteOut, error) { //nolint:unparam // UserAction signature requires returned parameter
	err := ua.authorizer.Authorize(
		ctx,
		user,
		authorization.NewAction(
			authorization.DeleteTagPermission,
			authorization.TagResource,
			in.ID,
		),
	)
	if err != nil {
		return TagUntagAllAndDeleteOut{}, authorization.NewUnauthorizedError(err)
	}

	err = ua.simpleActions.UntagAllAndDelete(ctx, in.ID)
	if err != nil {
		return TagUntagAllAndDeleteOut{}, fmt.Errorf("can't untag all and delete tag (ua): %w", err)
	}

	return TagUntagAllAndDeleteOut{}, nil
}
