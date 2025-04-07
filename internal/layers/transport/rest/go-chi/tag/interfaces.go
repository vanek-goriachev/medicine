package tag

import (
	"context"

	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	TagForceCreateInFromChi(in TagForceCreateIn) (tagUA.TagForceCreateIn, error)
	TagForceCreateOutToChi(out tagUA.TagForceCreateOut) TagForceCreateOut
}

type tagCreateUserAction interface {
	Act(ctx context.Context, user userModels.User, in tagUA.TagForceCreateIn) (tagUA.TagForceCreateOut, error)
}
