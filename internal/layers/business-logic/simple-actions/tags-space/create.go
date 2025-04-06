package tags_space

import (
	"context"
	"errors"
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
	pkgErrors "medicine/pkg/errors/db"
	userModels "medicine/pkg/user"
)

func (sa *SimpleActions) Create(
	ctx context.Context,
	user userModels.User,
	name string,
) (tagsSpaceModels.TagsSpace, error) {
	err := sa.checkExistence(ctx, user, name)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, err
	}

	tagsSpace, err := sa.build(user, name)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, err
	}

	err = sa.atomicActions.Create(ctx, tagsSpace)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't create tagsSpace: %w", err)
	}

	return tagsSpace, nil
}

func (sa *SimpleActions) checkExistence(ctx context.Context, user userModels.User, name string) error {
	identifier := customIdentifiers.UserIDAndNameIdentifier{UserID: user.ID, Name: name}

	_, err := sa.atomicActions.GetByUserIDAndName(ctx, identifier)
	if err == nil {
		return tagsSpaceModels.NewTagsSpaceAlreadyExistError(identifier)
	}

	spaceNotFoundErr := pkgErrors.NewDoesNotExistError(identifier)
	if !errors.Is(err, spaceNotFoundErr) {
		return fmt.Errorf("can't check tagsSpace existence: %w", err)
	}

	return nil
}

func (sa *SimpleActions) build(user userModels.User, name string) (tagsSpaceModels.TagsSpace, error) {
	id, err := sa.idGenerator.Generate()
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't generate an id: %w", err)
	}

	tagsSpace, err := sa.tagsSpaceFactory.New(id, user.ID, name, []tagModels.Tag{})
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't build tagsSpace: %w", err)
	}

	return tagsSpace, nil
}
