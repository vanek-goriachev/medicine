package tags_space

import (
	"context"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	TagsSpaceGetByIDInFromChi(in TagsSpaceGetByIDIn) (tagsSpaceUA.TagsSpaceGetByIDIn, error)
	TagsSpaceGetByIDOutToChi(out tagsSpaceUA.TagsSpaceGetByIDOut) TagsSpaceGetByIDOut

	TagsSpaceListByUserInFromChi(in TagsSpaceListByUserIn) (tagsSpaceUA.TagsSpaceListByUserIn, error)
	TagsSpaceListByUserOutToChi(out tagsSpaceUA.TagsSpaceListByUserOut) TagsSpaceListByUserOut

	TagsSpaceCreateInFromChi(in TagsSpaceCreateIn) (tagsSpaceUA.TagsSpaceCreateIn, error)
	TagsSpaceCreateOutToChi(out tagsSpaceUA.TagsSpaceCreateOut) TagsSpaceCreateOut

	TagsSpaceDeleteInFromChi(in TagsSpaceDeleteIn) (tagsSpaceUA.TagsSpaceDeleteIn, error)
	TagsSpaceDeleteOutToChi(_ tagsSpaceUA.TagsSpaceDeleteOut) TagsSpaceDeleteOut
}

type tagsSpaceCreateUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in tagsSpaceUA.TagsSpaceCreateIn,
	) (tagsSpaceUA.TagsSpaceCreateOut, error)
}

type tagsSpaceListByUserUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in tagsSpaceUA.TagsSpaceListByUserIn,
	) (tagsSpaceUA.TagsSpaceListByUserOut, error)
}

type tagsSpaceGetByIDUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in tagsSpaceUA.TagsSpaceGetByIDIn,
	) (tagsSpaceUA.TagsSpaceGetByIDOut, error)
}

type tagsSpaceDeleteUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in tagsSpaceUA.TagsSpaceDeleteIn,
	) (tagsSpaceUA.TagsSpaceDeleteOut, error)
}
