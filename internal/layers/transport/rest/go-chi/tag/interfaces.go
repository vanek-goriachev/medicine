package tag

import (
	"context"

	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	TagForceCreateInFromChi(in *TagForceCreateIn) (tagUA.TagForceCreateIn, error)
	TagForceCreateOutToChi(out *tagUA.TagForceCreateOut) TagForceCreateOut

	TagUntagAllAndDeleteInFromChi(in *TagUntagAllAndDeleteIn) (tagUA.TagUntagAllAndDeleteIn, error)
	TagUntagAllAndDeleteOutToChi(_ *tagUA.TagUntagAllAndDeleteOut) TagUntagAllAndDeleteOut
}

type tagCreateUserAction interface {
	Act(ctx context.Context, user userModels.User, in *tagUA.TagForceCreateIn) (tagUA.TagForceCreateOut, error)
}

type tagUntagAllAndDeleteUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in *tagUA.TagUntagAllAndDeleteIn,
	) (tagUA.TagUntagAllAndDeleteOut, error)
}
