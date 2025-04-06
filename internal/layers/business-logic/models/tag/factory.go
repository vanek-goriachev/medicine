package tag

import (
	"fmt"

	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

type Factory struct {
	validator validation.Validator[Tag]
}

func NewFactory(validator validation.Validator[Tag]) *Factory {
	return &Factory{
		validator: validator,
	}
}

func (f *Factory) New(id entityID.EntityID, spaceID entityID.EntityID, name string) (Tag, error) {
	tag := Tag{
		ID:          id,
		TagsSpaceID: spaceID,
		Name:        name,
	}

	validationError := f.validator.Validate(tag)
	if validationError != nil {
		return Tag{}, fmt.Errorf("validation error: %w", validationError)
	}

	return tag, nil
}
