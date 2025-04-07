package tags_space_test

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	tags_space "medicine/mocks/internal_/layers/business-logic/simple-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceListByUserSA(t *testing.T) {
	t.Parallel()

	user := generators.TestUser()
	expectedTagsSpaces := []tagsSpaceModels.TagsSpace{
		generators.GenerateTagsSpace(user.ID),
		generators.GenerateTagsSpace(user.ID),
		generators.GenerateTagsSpace(user.ID),
	}

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				ListByUserID(t.Context(), user.ID).
				Return(expectedTagsSpaces, nil)

			tagsSpaces, err := sa.ListByUser(t.Context(), user)

			assert.NoError(t, err)
			assert.Equal(t, expectedTagsSpaces, tagsSpaces)
		},
	)

	t.Run(
		"list by user id fail",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				ListByUserID(t.Context(), user.ID).
				Return(nil, assert.AnError)

			tagsSpaces, err := sa.ListByUser(t.Context(), user)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpaces)
		},
	)
}
