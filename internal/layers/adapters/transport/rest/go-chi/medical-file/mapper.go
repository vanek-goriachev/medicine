package medical_file

import (
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	medicalFileChi "medicine/internal/layers/transport/rest/go-chi/medical-file"
	entityID "medicine/pkg/entity-id"
)

type ChiMapper struct {
	entityIDMapper entityID.Mapper
}

func NewChiMapper(
	entityIDMapper entityID.Mapper,
) *ChiMapper {
	return &ChiMapper{
		entityIDMapper: entityIDMapper,
	}
}

func (m *ChiMapper) UploadedMedicalFileFromChi(
	uploadedFile medicalFileChi.UploadedMedicalFile,
) (medicalFileModels.UploadedMedicalFile, error) {
	return medicalFileModels.UploadedMedicalFile{
		Name: uploadedFile.Name,
		Data: uploadedFile.Data,
	}, nil
}

func (m *ChiMapper) MedicalFileInfoToChi(
	medicalFileInfo medicalFileModels.MedicalFileInfo,
) medicalFileChi.MedicalFileInfo {
	return medicalFileChi.MedicalFileInfo{
		ID:        medicalFileInfo.ID.String(),
		Extension: string(medicalFileInfo.Extension),
		Name:      medicalFileInfo.Name,
	}
}
