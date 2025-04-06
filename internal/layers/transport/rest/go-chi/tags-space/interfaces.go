package tags_space

import (
	"context"

	createUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	TagsSpaceGetByIDInFromChi(in TagsSpaceGetByIDIn) (createUA.TagsSpaceGetByIDIn, error)
	TagsSpaceGetByIDOutToChi(out createUA.TagsSpaceGetByIDOut) TagsSpaceGetByIDOut

	TagsSpaceListByUserInFromChi(in TagsSpaceListByUserIn) (createUA.TagsSpaceListByUserIn, error)
	TagsSpaceListByUserOutToChi(out createUA.TagsSpaceListByUserOut) TagsSpaceListByUserOut

	TagsSpaceCreateInFromChi(in TagsSpaceCreateIn) (createUA.TagsSpaceCreateIn, error)
	TagsSpaceCreateOutToChi(out createUA.TagsSpaceCreateOut) TagsSpaceCreateOut
}

type createTagsSpaceUserAction interface {
	Act(ctx context.Context, user userModels.User, in createUA.TagsSpaceCreateIn) (createUA.TagsSpaceCreateOut, error)
}

type tagsSpaceListByUserUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in createUA.TagsSpaceListByUserIn,
	) (createUA.TagsSpaceListByUserOut, error)
}

type tagsSpaceGetByIDUserAction interface {
	Act(ctx context.Context, user userModels.User, in createUA.TagsSpaceGetByIDIn) (createUA.TagsSpaceGetByIDOut, error)
}
