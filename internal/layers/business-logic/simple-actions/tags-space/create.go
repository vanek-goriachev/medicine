package tags_space

import (
	"context"
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

func (sa *SimpleActions) Create(
	ctx context.Context,
	_ userModels.User,
	name string,
) (tagsSpaceModels.TagsSpace, error) {
	tagsSpace, err := sa.build(name)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, err
	}

	err = sa.tagsSpaceAtomicActions.Create(ctx, tagsSpace)
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't create tagsSpace: %w", err)
	}

	// TODO: write to IAM that new tagsSpace created and owner by user

	return tagsSpace, nil
}

func (sa *SimpleActions) build(name string) (tagsSpaceModels.TagsSpace, error) {
	id, err := sa.idGenerator.Generate()
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't generate an id: %w", err)
	}

	tagsSpace, err := sa.tagsSpaceFactory.New(id, name, []tagModels.Tag{})
	if err != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("can't build tagsSpace: %w", err)
	}

	return tagsSpace, nil
}
