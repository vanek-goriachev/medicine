package visit_record

import (
	"fmt"
	"time"

	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

type Factory struct {
	validator validation.Validator[VisitRecord]
}

func NewFactory(validator validation.Validator[VisitRecord]) *Factory {
	return &Factory{
		validator: validator,
	}
}

func (f *Factory) New(
	id entityID.EntityID,
	name string,
	datetime time.Time,
) (VisitRecord, error) {
	visitRecord := VisitRecord{
		ID:       id,
		Name:     name,
		Datetime: datetime,
	}

	validationError := f.validator.Validate(visitRecord)
	if validationError != nil {
		return VisitRecord{}, fmt.Errorf("validation error: %w", validationError)
	}

	return visitRecord, nil
}
