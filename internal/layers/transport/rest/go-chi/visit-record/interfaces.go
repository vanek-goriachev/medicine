package visit_record

import (
	"context"

	visitRecordUA "medicine/internal/layers/business-logic/user-actions/visit-record"
	userModels "medicine/pkg/user"
)

type userActionsMapper interface {
	VisitRecordCreateInFromChi(in *VisitRecordCreateIn) (visitRecordUA.VisitRecordCreateIn, error)
	VisitRecordCreateOutToChi(out *visitRecordUA.VisitRecordCreateOut) VisitRecordCreateOut
}

type visitRecordCreateUserAction interface {
	Act(
		ctx context.Context,
		user userModels.User,
		in *visitRecordUA.VisitRecordCreateIn,
	) (visitRecordUA.VisitRecordCreateOut, error)
}
