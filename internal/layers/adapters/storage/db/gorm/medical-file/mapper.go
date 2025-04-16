package medical_file

import (
	"github.com/google/uuid"

	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
)

type GORMMapper struct{}

func NewGORMMapper() *GORMMapper {
	return &GORMMapper{}
}

func (*GORMMapper) ToGORM(medicalFile medicalFileModels.MedicalFile) (
	gormModels.MedicalFileInfo,
	gormModels.MedicalFileData,
) {
	info := gormModels.MedicalFileInfo{
		Name: medicalFile.Name,
		ID:   uuid.UUID(medicalFile.ID),
	}

	data := gormModels.MedicalFileData{
		ID:   uuid.UUID(medicalFile.ID),
		Data: gormModels.DataType(medicalFile.Data),
	}

	return info, data
}
