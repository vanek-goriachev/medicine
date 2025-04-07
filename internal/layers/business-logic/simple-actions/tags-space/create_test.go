package tags_space_test

import (
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	customIdentifiers "medicine/internal/layers/business-logic/models/tags-space/identifiers"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	tags_space "medicine/mocks/internal_/layers/business-logic/simple-actions/tags-space"
	entityID "medicine/pkg/entity-id"
	pkgErrors "medicine/pkg/errors/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceCreateSA(t *testing.T) {
	t.Parallel()

	user := generators.TestUser()
	name := "tags-space"
	identifier := customIdentifiers.UserIDAndNameIdentifier{
		UserID: user.ID,
		Name:   name,
	}
	spaceNotFoundErr := pkgErrors.NewDoesNotExistError(identifier)
	tagsSpaceID := generators.GenerateEntityID()
	expectedTagsSpace := tagsSpaceModels.TagsSpace{
		ID:     tagsSpaceID,
		UserID: user.ID,
		Name:   name,
		Tags:   []tagModels.Tag{},
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
				GetByUserIDAndName(t.Context(), identifier).
				Return(tagsSpaceModels.TagsSpace{}, spaceNotFoundErr)
			idGenerator.EXPECT().Generate().Return(tagsSpaceID, nil)
			tagsSpaceFactory.EXPECT().
				New(tagsSpaceID, user.ID, name, []tagModels.Tag{}).
				Return(expectedTagsSpace, nil)
			atomicActions.EXPECT().Create(t.Context(), expectedTagsSpace).Return(nil)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.NoError(t, err)
			assert.Equal(t, expectedTagsSpace, tagsSpace)
		},
	)

	t.Run(
		"create fail",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				GetByUserIDAndName(t.Context(), identifier).
				Return(tagsSpaceModels.TagsSpace{}, spaceNotFoundErr)
			idGenerator.EXPECT().Generate().Return(tagsSpaceID, nil)
			tagsSpaceFactory.EXPECT().
				New(tagsSpaceID, user.ID, name, []tagModels.Tag{}).
				Return(expectedTagsSpace, nil)
			atomicActions.EXPECT().Create(t.Context(), expectedTagsSpace).Return(assert.AnError)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)

	t.Run(
		"factory fail",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				GetByUserIDAndName(t.Context(), identifier).
				Return(tagsSpaceModels.TagsSpace{}, spaceNotFoundErr)
			idGenerator.EXPECT().Generate().Return(tagsSpaceID, nil)
			tagsSpaceFactory.EXPECT().
				New(tagsSpaceID, user.ID, name, []tagModels.Tag{}).
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

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				GetByUserIDAndName(t.Context(), identifier).
				Return(tagsSpaceModels.TagsSpace{}, spaceNotFoundErr)
			idGenerator.EXPECT().Generate().Return(entityID.EntityID{}, assert.AnError)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)

	t.Run(
		"get tags space fail",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				GetByUserIDAndName(t.Context(), identifier).
				Return(tagsSpaceModels.TagsSpace{}, assert.AnError)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpace)
		},
	)

	t.Run(
		"tags space already exists",
		func(t *testing.T) {
			t.Parallel()

			idGenerator := tags_space.NewEntityIDGenerator(t)
			tagsSpaceFactory := tags_space.NewTagsSpaceFactory(t)
			atomicActions := tags_space.NewAtomicActions(t)
			sa := tagsSpaceSA.NewSimpleActions(idGenerator, tagsSpaceFactory, atomicActions)

			atomicActions.EXPECT().
				GetByUserIDAndName(t.Context(), identifier).
				Return(expectedTagsSpace, nil)

			tagsSpace, err := sa.Create(t.Context(), user, name)

			assert.ErrorIs(t, err, tagsSpaceModels.NewTagsSpaceAlreadyExistError(identifier))
			assert.Zero(t, tagsSpace)
		},
	)

}
