package dto

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
)

type UploadedMedicalFile struct {
	Name string
	Data []byte
}

func ParseUploadedMedicalFiles(formFiles []*multipart.FileHeader) ([]UploadedMedicalFile, error) {
	var parseErrors []error
	uploadedMedicalFiles := make([]UploadedMedicalFile, len(formFiles))

	for i, formFile := range formFiles {
		var err error

		uploadedMedicalFiles[i], err = ParseUploadedMedicalFile(formFile)
		if err != nil {
			parseErrors = append(parseErrors, err)
		}
	}

	if len(parseErrors) > 0 {
		return nil, errors.Join(parseErrors...)
	}

	return uploadedMedicalFiles, nil
}

func ParseUploadedMedicalFile(formFile *multipart.FileHeader) (UploadedMedicalFile, error) {
	file, err := formFile.Open()
	if err != nil {
		return UploadedMedicalFile{}, fmt.Errorf(
			"failed to open uploaded file with name %s: %w",
			formFile.Filename,
			err,
		)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return UploadedMedicalFile{}, fmt.Errorf(
			"failed to read uploaded file with name = %s: %w",
			formFile.Filename,
			err,
		)
	}

	return UploadedMedicalFile{
		Name: formFile.Filename,
		Data: data,
	}, nil
}
