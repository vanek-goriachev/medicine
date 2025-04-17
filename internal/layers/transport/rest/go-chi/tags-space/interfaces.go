package tags_space

import (
	"context"

	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	TagsSpaceGetByIDInFromChi(in *TagsSpaceGetByIDIn) (tagsSpaceUA.TagsSpaceGetByIDIn, error)
	TagsSpaceGetByIDOutToChi(out *tagsSpaceUA.TagsSpaceGetByIDOut) TagsSpaceGetByIDOut

	TagsSpaceListAllAvailableInFromChi(in *TagsSpaceListAllAvailableIn) (tagsSpaceUA.TagsSpaceListAllAvailableIn, error)
	TagsSpaceListAllAvailableOutToChi(out *tagsSpaceUA.TagsSpaceListAllAvailableOut) TagsSpaceListAllAvailableOut

	TagsSpaceCreateInFromChi(in *TagsSpaceCreateIn) (tagsSpaceUA.TagsSpaceCreateIn, error)
	TagsSpaceCreateOutToChi(out *tagsSpaceUA.TagsSpaceCreateOut) TagsSpaceCreateOut

	TagsSpaceDeleteInFromChi(in *TagsSpaceDeleteIn) (tagsSpaceUA.TagsSpaceDeleteIn, error)
	TagsSpaceDeleteOutToChi(_ *tagsSpaceUA.TagsSpaceDeleteOut) TagsSpaceDeleteOut
}

type tagsSpaceCreateUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in *tagsSpaceUA.TagsSpaceCreateIn,
	) (tagsSpaceUA.TagsSpaceCreateOut, error)
}

type tagsSpaceListAllAvailableUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in *tagsSpaceUA.TagsSpaceListAllAvailableIn,
	) (tagsSpaceUA.TagsSpaceListAllAvailableOut, error)
}

type tagsSpaceGetByIDUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in *tagsSpaceUA.TagsSpaceGetByIDIn,
	) (tagsSpaceUA.TagsSpaceGetByIDOut, error)
}

type tagsSpaceDeleteUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in *tagsSpaceUA.TagsSpaceDeleteIn,
	) (tagsSpaceUA.TagsSpaceDeleteOut, error)
}
