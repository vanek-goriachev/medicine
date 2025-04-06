package tag

import (
	"errors"
	"fmt"

	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

const MaxTagNameLength = 32

var (
	ErrTagNameRequired = errors.New("tag name is required")
	ErrTagNameTooLong  = fmt.Errorf("tag name should be less than %d characters", MaxTagNameLength)
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

func (v *Validator) Validate(tag Tag) error {
	idErr := v.idValidator.Validate(tag.ID)
	tagsSpaceIDErr := v.idValidator.Validate(tag.TagsSpaceID)

	nameErr := v.validateName(tag.Name)

	return validation.NewValidationError(
		errors.Join(idErr, tagsSpaceIDErr, nameErr),
	)
}

func (*Validator) validateName(name string) error {
	if name == "" {
		return ErrTagNameRequired
	}

	if len(name) > MaxTagNameLength {
		return ErrTagNameTooLong
	}

	return nil
}
