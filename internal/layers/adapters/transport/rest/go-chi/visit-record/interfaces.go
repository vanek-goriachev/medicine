package visit_record

import (
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	medicalFileChi "medicine/internal/layers/transport/rest/go-chi/medical-file"
	visitRecordChi "medicine/internal/layers/transport/rest/go-chi/visit-record"
)

type medicalFileChiMapper interface {
	UploadedMedicalFileFromChi(
		uploadedFile medicalFileChi.UploadedMedicalFile,
	) (medicalFileModels.UploadedMedicalFile, error)

	MedicalFileInfoToChi(
		medicalFileInfo medicalFileModels.MedicalFileInfo,
	) medicalFileChi.MedicalFileInfo
}

// visitRecordChiMapper implemented by ChiMapper.
type visitRecordChiMapper interface {
	ToChi(visitRecord visitRecordModels.VisitRecord) visitRecordChi.VisitRecord

	LinkedEntitiesToChi(
		visitRecordLinkedEntities visitRecordModels.VisitRecordLinkedEntities,
	) visitRecordChi.VisitRecordLinkedEntities
}
