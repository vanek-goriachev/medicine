package tags_space

import (
	"context"

	createUA "medicine/internal/layers/business-logic/user-actions/tags-space"
)

type userActionsMapper interface {
	CreateTagsSpaceInFromChi(in CreateTagsSpaceIn) (createUA.CreateTagsSpaceIn, error)
	CreateTagsSpaceOutToChi(out createUA.CreateTagsSpaceOut) CreateTagsSpaceOut
}

type createTagsSpaceUserAction interface {
	Act(ctx context.Context, in createUA.CreateTagsSpaceIn) (createUA.CreateTagsSpaceOut, error)
}
