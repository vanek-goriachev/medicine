package tags_space

import (
	"context"

	createUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	CreateTagsSpaceInFromChi(in CreateTagsSpaceIn) createUA.CreateTagsSpaceIn
	CreateTagsSpaceOutToChi(out createUA.CreateTagsSpaceOut) CreateTagsSpaceOut
}

type createTagsSpaceUserAction interface {
	Act(ctx context.Context, user userModels.User, in createUA.CreateTagsSpaceIn) (createUA.CreateTagsSpaceOut, error)
}
