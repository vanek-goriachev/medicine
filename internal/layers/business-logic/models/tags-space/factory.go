package tags_space

import (
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	"medicine/internal/tooling/validation"
	entityID "medicine/pkg/entity-id"
)

type Factory struct {
	validator validation.Validator[TagsSpace]
}

func NewFactory(validator validation.Validator[TagsSpace]) *Factory {
	return &Factory{
		validator: validator,
	}
}

func (f *Factory) New(
	id entityID.EntityID,
	name string,
	tags []tagModels.Tag,
) (TagsSpace, error) {
	tagsSpace := TagsSpace{
		ID:   id,
		Name: name,
		Tags: tags,
	}

	validationError := f.validator.Validate(tagsSpace)
	if validationError != nil {
		return TagsSpace{}, fmt.Errorf("validation error: %w", validationError)
	}

	return tagsSpace, nil
}
