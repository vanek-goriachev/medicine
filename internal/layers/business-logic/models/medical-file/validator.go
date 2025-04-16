package medical_file

import (
	"errors"
	"fmt"

	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

const MaxMedicalFileNameLength = 64
const MaxMedicalFileDataLength = 1024 * 1024 * 1024 // 1GB

var (
	ErrMedicalFileDataRequired      = errors.New("medicalFile data is required")
	ErrMedicalFileDataTooLarge      = fmt.Errorf("medicalFile data should be less than %d bytes", MaxMedicalFileDataLength)
	ErrMedicalFileExtensionRequired = errors.New("medicalFile extension is required")
	ErrMedicalFileNameRequired      = errors.New("medicalFile name is required")
	ErrMedicalFileNameTooLong       = fmt.Errorf("medicalFile name should be less than %d characters", MaxMedicalFileNameLength)
)

type Validator struct {
	idValidator validation.Validator[entityID.EntityID]
}

func NewValidator(
	idValidator validation.Validator[entityID.EntityID],
) *Validator {
	return &Validator{
		idValidator: idValidator,
	}
}

func (v *Validator) Validate(medicalFile MedicalFile) error {
	infoErr := v.validateInfo(medicalFile.MedicalFileInfo)
	dataErr := v.validateData(medicalFile.MedicalFileData)

	validationErrors := errors.Join(infoErr, dataErr)
	if validationErrors != nil {
		return validation.NewValidationError(validationErrors)
	}

	return nil
}

func (v *Validator) validateInfo(info MedicalFileInfo) error {
	idErr := v.idValidator.Validate(info.ID)
	nameErr := v.validateName(info.Name)
	extensionErr := v.validateExtension(info.Extension)

	validationErrors := errors.Join(idErr, nameErr, extensionErr)
	if validationErrors != nil {
		return validation.NewValidationError(validationErrors)
	}

	return nil
}

func (v *Validator) validateName(name string) error {
	if name == "" {
		return ErrMedicalFileNameRequired
	}

	if len(name) > MaxMedicalFileNameLength {
		return ErrMedicalFileNameTooLong
	}

	return nil
}

func (v *Validator) validateExtension(extension Extension) error {
	if extension == EmptyFileException {
		return ErrMedicalFileExtensionRequired
	}

	return nil
}

func (v *Validator) validateData(data MedicalFileData) error {
	if len(*data.Data) == 0 {
		return ErrMedicalFileDataRequired
	}

	if len(*data.Data) > MaxMedicalFileDataLength {
		return ErrMedicalFileDataTooLarge
	}

	return nil
}
