package medical_file

import (
	"errors"
	"fmt"

	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	medicalFileChi "medicine/internal/layers/transport/rest/go-chi/medical-file/dto"
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

func (*ChiMapper) UploadedMedicalFileFromChi(
	uploadedFile medicalFileChi.UploadedMedicalFile,
) (medicalFileModels.UploadedMedicalFile, error) { //nolint:unparam // Required signature
	return medicalFileModels.UploadedMedicalFile{
		Name: uploadedFile.Name,
		Data: uploadedFile.Data,
	}, nil
}

func (m *ChiMapper) MultipleUploadedMedicalFileFromChi(
	uploadedFiles []medicalFileChi.UploadedMedicalFile,
) ([]medicalFileModels.UploadedMedicalFile, error) {
	var convertErrors []error
	uploadedMedicalFiles := make([]medicalFileModels.UploadedMedicalFile, len(uploadedFiles))

	for i, uploadedFile := range uploadedFiles {
		var err error

		uploadedMedicalFiles[i], err = m.UploadedMedicalFileFromChi(uploadedFile)
		if err != nil {
			convertErrors = append(convertErrors, err)
			continue
		}
	}

	if len(convertErrors) > 0 {
		return nil, fmt.Errorf("couldn't convert uploaded medical files: %w", errors.Join(convertErrors...))
	}

	return uploadedMedicalFiles, nil
}

func (*ChiMapper) MedicalFileInfoToChi(
	medicalFileInfo medicalFileModels.MedicalFileInfo,
) medicalFileChi.MedicalFileInfo {
	return medicalFileChi.MedicalFileInfo{
		ID:        medicalFileInfo.ID.String(),
		Extension: string(medicalFileInfo.Extension),
		Name:      medicalFileInfo.Name,
	}
}
