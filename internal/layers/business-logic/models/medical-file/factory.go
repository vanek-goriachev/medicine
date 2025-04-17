package medical_file

import (
	"fmt"
	"path"

	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

type Factory struct {
	validator validation.Validator[MedicalFile]
}

func NewFactory(validator validation.Validator[MedicalFile]) *Factory {
	return &Factory{
		validator: validator,
	}
}

func (f *Factory) New(
	id entityID.EntityID,
	name string,
	data DataType,
) (MedicalFile, error) {
	medicalFile := MedicalFile{
		MedicalFileInfo: MedicalFileInfo{
			ID:        id,
			Extension: Extension(path.Ext(name)),
			Name:      name,
		},
		MedicalFileData: MedicalFileData{
			Data: data,
		},
	}

	validationError := f.validator.Validate(medicalFile)
	if validationError != nil {
		return MedicalFile{}, fmt.Errorf("validation error: %w", validationError)
	}

	return medicalFile, nil
}
