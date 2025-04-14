package tags_space_test

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	tags_space "medicine/mocks/internal_/layers/business-logic/simple-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceGetByIDSA(t *testing.T) {
	t.Parallel()

	user := generators.TestUser()
	tagsSpaceID := generators.GenerateEntityID()
	expectedTagsSpace := generators.GenerateTagsSpace(user.ID)

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			tagsSpaceAtomicActions := tags_space.NewTagsSpaceAtomicActions(t)
			tagAtomicActions := tags_space.NewTagAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, tagAtomicActions, tagsSpaceAtomicActions)

			tagsSpaceAtomicActions.EXPECT().
				GetByID(t.Context(), tagsSpaceID).
				Return(expectedTagsSpace, nil)

			tagsSpace, err := sa.GetByID(t.Context(), tagsSpaceID)

			assert.NoError(t, err)
			assert.Equal(t, expectedTagsSpace, tagsSpace)
		},
	)

	t.Run(
		"get by id fail",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			tagsSpaceAtomicActions := tags_space.NewTagsSpaceAtomicActions(t)
			tagAtomicActions := tags_space.NewTagAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, tagAtomicActions, tagsSpaceAtomicActions)

			tagsSpaceAtomicActions.EXPECT().
				GetByID(t.Context(), tagsSpaceID).
				Return(tagsSpaceModels.TagsSpace{}, assert.AnError)

			tagsSpace, err := sa.GetByID(t.Context(), tagsSpaceID)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)
}
