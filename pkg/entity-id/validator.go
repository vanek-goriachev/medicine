package entity_id

import (
	"github.com/google/uuid"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (*Validator) Validate(id EntityID) error {
	err := uuid.Validate(id.String())

	if err != nil {
		return NewInvalidEntityIDError(id.String())
	}

	return nil
}
