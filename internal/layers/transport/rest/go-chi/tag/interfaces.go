package tag

import (
	"context"

	tagUA "medicine/internal/layers/business-logic/user-actions/tag"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	//TagGetByIDInFromChi(in TagGetByIDIn) (tagUA.TagGetByIDIn, error)
	//TagGetByIDOutToChi(out tagUA.TagGetByIDOut) TagGetByIDOut

	TagForceCreateInFromChi(in TagForceCreateIn) (tagUA.TagForceCreateIn, error)
	TagForceCreateOutToChi(out tagUA.TagForceCreateOut) TagForceCreateOut
}

type tagCreateUserAction interface {
	Act(ctx context.Context, user userModels.User, in tagUA.TagForceCreateIn) (tagUA.TagForceCreateOut, error)
}

//type TagGetByIDUserAction interface {
//	Act(ctx context.Context, user userModels.User, in tagUA.TagGetByIDIn) (tagUA.TagGetByIDOut, error)
//}
