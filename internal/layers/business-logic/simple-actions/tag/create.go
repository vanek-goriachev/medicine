package tag

import (
	"context"
	"errors"
	"fmt"
	customIdentifiers "medicine/internal/layers/business-logic/models/tag/identifiers"
	entityID "medicine/pkg/entity-id"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	pkgErrors "medicine/pkg/errors/db"
)

func (sa *SimpleActions) Create(
	ctx context.Context,
	name string,
	tagsSpaceID entityID.EntityID,
) (tagModels.Tag, error) {
	err := sa.checkExistence(ctx, name, tagsSpaceID)
	if err != nil {
		return tagModels.Tag{}, err
	}

	tag, err := sa.build(name, tagsSpaceID)
	if err != nil {
		return tagModels.Tag{}, err
	}

	err = sa.atomicActions.Create(ctx, tag)
	if err != nil {
		return tagModels.Tag{}, fmt.Errorf("can't create tag: %w", err)
	}

	return tag, nil
}

func (sa *SimpleActions) checkExistence(ctx context.Context, name string, tagsSpaceID entityID.EntityID) error {
	identifier := customIdentifiers.TagsSpaceIDAndNameIdentifier{TagsSpaceID: tagsSpaceID, Name: name}

	_, err := sa.atomicActions.GetByTagsSpaceIDAndName(ctx, identifier)
	if err == nil {
		return tagModels.NewTagAlreadyExistError(identifier)
	}

	tagNotFoundErr := pkgErrors.NewDoesNotExistError(identifier)
	if !errors.Is(err, tagNotFoundErr) {
		return fmt.Errorf("can't check tag existence: %w", err)
	}

	return nil
}

func (sa *SimpleActions) build(name string, tagsSpaceID entityID.EntityID) (tagModels.Tag, error) {
	id, err := sa.idGenerator.Generate()
	if err != nil {
		return tagModels.Tag{}, fmt.Errorf("can't generate an id: %w", err)
	}

	tag, err := sa.tagFactory.New(id, tagsSpaceID, name)
	if err != nil {
		return tagModels.Tag{}, fmt.Errorf("can't build tag: %w", err)
	}

	return tag, nil
}
