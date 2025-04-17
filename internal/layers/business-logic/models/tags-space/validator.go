package tags_space

import (
	"errors"
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

const MaxTagsSpaceNameLength = 32

var (
	ErrTagsSpaceNameRequired = errors.New("tags space name is required")
	ErrTagsSpaceNameTooLong  = fmt.Errorf("tags space name should be less than %d characters", MaxTagsSpaceNameLength)
)

type Validator struct {
	idValidator  validation.Validator[entityID.EntityID]
	tagValidator validation.Validator[tagModels.Tag]
}

func NewValidator(
	idValidator validation.Validator[entityID.EntityID],
	tagValidator validation.Validator[tagModels.Tag],
) *Validator {
	return &Validator{
		idValidator:  idValidator,
		tagValidator: tagValidator,
	}
}

func (v *Validator) Validate(tagsSpace TagsSpace) error {
	idErr := v.idValidator.Validate(tagsSpace.ID)

	nameErr := v.validateName(tagsSpace.Name)

	tagsErrorsArr := make([]error, len(tagsSpace.Tags))
	for _, tag := range tagsSpace.Tags {
		tagsErrorsArr = append(tagsErrorsArr, v.tagValidator.Validate(tag))
	}

	tagErrors := errors.Join(tagsErrorsArr...)

	validationErrors := errors.Join(idErr, nameErr, tagErrors)
	if validationErrors != nil {
		return validation.NewValidationError(validationErrors)
	}

	return nil
}

func (*Validator) validateName(name string) error {
	if name == "" {
		return ErrTagsSpaceNameRequired
	}

	if len(name) > MaxTagsSpaceNameLength {
		return ErrTagsSpaceNameTooLong
	}

	return nil
}
