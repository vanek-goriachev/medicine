package tags_space_test

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	"medicine/internal/tooling/tests"
	tags_space "medicine/mocks/internal_/layers/business-logic/simple-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceGetByIDSA(t *testing.T) {
	t.Parallel()

	user := tests.TestUser()
	tagsSpaceID := tests.GenerateEntityID()
	expectedTagsSpace := tagsSpaceModels.TagsSpace{
		ID:     tagsSpaceID,
		UserID: user.ID,
		Name:   "some-name",
		Tags: []tagModels.Tag{
			{
				ID:          tests.GenerateEntityID(),
				TagsSpaceID: tagsSpaceID,
				Name:        "tags1",
			},
			{
				ID:          tests.GenerateEntityID(),
				TagsSpaceID: tagsSpaceID,
				Name:        "tags2",
			},
		},
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
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				GetByID(t.Context(), tagsSpaceID).
				Return(tagsSpaceModels.TagsSpace{}, assert.AnError)

			tagsSpace, err := sa.GetByID(t.Context(), tagsSpaceID)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)
}
