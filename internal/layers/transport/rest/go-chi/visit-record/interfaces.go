package visit_record

import (
	ua "medicine/pkg/user-action"

	visitRecordUA "medicine/internal/layers/business-logic/user-actions/visit-record"
)

type userActionsMapper interface {
	VisitRecordCreateInFromChi(in *VisitRecordCreateIn) (visitRecordUA.VisitRecordCreateIn, error)
	VisitRecordCreateOutToChi(out *visitRecordUA.VisitRecordCreateOut) VisitRecordCreateOut

	VisitRecordAttachMedicalFilesInFromChi(in *VisitRecordAttachMedicalFilesIn) (visitRecordUA.VisitRecordAttachMedicalFilesIn, error)
	VisitRecordAttachMedicalFilesOutToChi(out *visitRecordUA.VisitRecordAttachMedicalFilesOut) VisitRecordAttachMedicalFilesOut
}

type visitRecordCreateUserAction = ua.UserAction[visitRecordUA.VisitRecordCreateIn, visitRecordUA.VisitRecordCreateOut]
type visitRecordAttachMedicalFilesUserAction = ua.UserAction[
	visitRecordUA.VisitRecordAttachMedicalFilesIn,
	visitRecordUA.VisitRecordAttachMedicalFilesOut,
]
