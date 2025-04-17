package tags_space_test

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	"medicine/mocks/internal_/layers/business-logic/authorization"
	tags_space "medicine/mocks/internal_/layers/business-logic/simple-actions/tags-space"
	entityID "medicine/pkg/entity-id"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceCreateSA(t *testing.T) {
	t.Parallel()

	user := generators.TestUser()
	name := "tags-space"
	tagsSpaceID := generators.GenerateEntityID()
	expectedTagsSpace := tagsSpaceModels.TagsSpace{
		ID:   tagsSpaceID,
		Name: name,
		Tags: []tagModels.Tag{},
	}

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorization.NewAuthorizer(t)
			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			tagsSpaceAtomicActions := tags_space.NewTagsSpaceAtomicActions(t)
			tagAtomicActions := tags_space.NewTagAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(
				authorizer,
				idGenerator,
				tagsSpaceFactory,
				tagAtomicActions,
				tagsSpaceAtomicActions,
			)

			idGenerator.EXPECT().Generate().Return(tagsSpaceID, nil)
			tagsSpaceFactory.EXPECT().
				New(tagsSpaceID, name, []tagModels.Tag{}).
				Return(expectedTagsSpace, nil)
			tagsSpaceAtomicActions.EXPECT().Create(t.Context(), expectedTagsSpace).Return(nil)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.NoError(t, err)
			assert.Equal(t, expectedTagsSpace, tagsSpace)
		},
	)

	t.Run(
		"create fail",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorization.NewAuthorizer(t)
			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			tagsSpaceAtomicActions := tags_space.NewTagsSpaceAtomicActions(t)
			tagAtomicActions := tags_space.NewTagAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(
				authorizer,
				idGenerator,
				tagsSpaceFactory,
				tagAtomicActions,
				tagsSpaceAtomicActions,
			)

			idGenerator.EXPECT().Generate().Return(tagsSpaceID, nil)
			tagsSpaceFactory.EXPECT().
				New(tagsSpaceID, name, []tagModels.Tag{}).
				Return(expectedTagsSpace, nil)
			tagsSpaceAtomicActions.EXPECT().Create(t.Context(), expectedTagsSpace).Return(assert.AnError)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)

	t.Run(
		"factory fail",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorization.NewAuthorizer(t)
			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			tagsSpaceAtomicActions := tags_space.NewTagsSpaceAtomicActions(t)
			tagAtomicActions := tags_space.NewTagAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(
				authorizer,
				idGenerator,
				tagsSpaceFactory,
				tagAtomicActions,
				tagsSpaceAtomicActions,
			)

			idGenerator.EXPECT().Generate().Return(tagsSpaceID, nil)
			tagsSpaceFactory.EXPECT().
				New(tagsSpaceID, name, []tagModels.Tag{}).
				Return(tagsSpaceModels.TagsSpace{}, assert.AnError)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)

	t.Run(
		"id generator fail",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorization.NewAuthorizer(t)
			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			tagsSpaceAtomicActions := tags_space.NewTagsSpaceAtomicActions(t)
			tagAtomicActions := tags_space.NewTagAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(
				authorizer,
				idGenerator,
				tagsSpaceFactory,
				tagAtomicActions,
				tagsSpaceAtomicActions,
			)

			idGenerator.EXPECT().Generate().Return(entityID.EntityID{}, assert.AnError)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)
}
