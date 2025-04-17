package medical_file

import (
	"context"
	"errors"
	"fmt"

	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
)

func (sa *SimpleActions) CreateMultiple(
	ctx context.Context,
	uploadedFiles []medicalFileModels.UploadedMedicalFile,
) ([]medicalFileModels.MedicalFile, error) {
	files, err := sa.buildFiles(uploadedFiles)
	if err != nil {
		return nil, err
	}

	err = sa.createMany(ctx, files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (sa *SimpleActions) buildFiles(
	uploadedFiles []medicalFileModels.UploadedMedicalFile,
) ([]medicalFileModels.MedicalFile, error) {
	var buildFilesErrors []error
	files := make([]medicalFileModels.MedicalFile, len(uploadedFiles))

	for i, uploadedFile := range uploadedFiles {
		var err error

		files[i], err = sa.buildFile(uploadedFile)
		if err != nil {
			buildFilesErrors = append(buildFilesErrors, err)
		}
	}

	if len(buildFilesErrors) > 0 {
		return nil, fmt.Errorf(
			"failed to build medical files: %w",
			errors.Join(buildFilesErrors...),
		)
	}

	return files, nil
}

func (sa *SimpleActions) buildFile(
	uploadedFile medicalFileModels.UploadedMedicalFile,
) (medicalFileModels.MedicalFile, error) {
	id, err := sa.idGenerator.Generate()
	if err != nil {
		return medicalFileModels.MedicalFile{}, fmt.Errorf("failed to generate id: %w", err)
	}

	medicalFile, err := sa.medicalFileFactory.New(
		id,
		uploadedFile.Name,
		uploadedFile.Data,
	)
	if err != nil {
		return medicalFile, fmt.Errorf("failed to build medical file with name='%s': %w", uploadedFile.Name, err)
	}

	return medicalFile, nil
}

func (sa *SimpleActions) createMany(
	ctx context.Context,
	files []medicalFileModels.MedicalFile,
) error {
	// TODO replace with multiple files create
	for _, file := range files {
		err := sa.atomicActions.Create(ctx, file)
		if err != nil {
			return fmt.Errorf("failed to create medical file with name='%s': %w", file.Name, err)
		}
	}

	return nil
}
