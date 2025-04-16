package visit_record

import (
	"errors"
	"fmt"
	"time"

	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

const MaxVisitRecordNameLength = 255

var (
	ErrVisitRecordDatetimeRequired = errors.New("visitRecord datetime is required")
	ErrVisitRecordNameRequired     = errors.New("visitRecord name is required")
	ErrVisitRecordNameTooLong      = fmt.Errorf("visitRecord namehould be less than %d characters", MaxVisitRecordNameLength)
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

func (v *Validator) Validate(visitRecord VisitRecord) error {
	idErr := v.idValidator.Validate(visitRecord.ID)

	nameErr := v.validateName(visitRecord.Name)

	datetimeErr := v.validateDatetime(visitRecord.Datetime)

	validationErrors := errors.Join(idErr, nameErr, datetimeErr)
	if validationErrors != nil {
		return validation.NewValidationError(validationErrors)
	}

	return nil
}

func (*Validator) validateName(name string) error {
	if name == "" {
		return ErrVisitRecordNameRequired
	}

	if len(name) > MaxVisitRecordNameLength {
		return ErrVisitRecordNameTooLong
	}

	return nil
}

func (*Validator) validateDatetime(datetime time.Time) error {
	if datetime.IsZero() {
		return ErrVisitRecordDatetimeRequired
	}

	return nil
}
