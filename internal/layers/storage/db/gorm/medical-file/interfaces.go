package medical_file

import (
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
)

type medicalFileGORMMapper interface {
	ToGORM(medicalFile medicalFileModels.MedicalFile) (
		gormModels.MedicalFileInfo,
		gormModels.MedicalFileData,
	)
	// FromGORM(
	//	dbMedicalFileInfo gormModels.MedicalFileInfo,
	//	dbMedicalFileData gormModels.MedicalFileData,
	// ) medicalFileModels.MedicalFile

	// InfoFromGORM(dbMedicalFileInfo gormModels.MedicalFileInfo) medicalFileModels.MedicalFileInfo
	// MultipleInfoFromGORM(dbMedicalFileInfos []gormModels.MedicalFileInfo) []medicalFileModels.MedicalFile
	// InfoToGORM(medicalFileInfo medicalFileModels.MedicalFileInfo) gormModels.MedicalFileInfo
	// MultipleInfoToGORM(medicalFileInfos []medicalFileModels.MedicalFile) []gormModels.MedicalFileInfo
	//
	// DataFromGORM(dbMedicalFileData gormModels.MedicalFileData) medicalFileModels.MedicalFileData
	// DataToGORM(medicalFileData medicalFileModels.MedicalFileData) gormModels.MedicalFileData
}
