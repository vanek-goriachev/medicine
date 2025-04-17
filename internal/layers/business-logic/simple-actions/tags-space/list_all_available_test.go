package tags_space_test

import (
	"context"
	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceSA "medicine/internal/layers/business-logic/simple-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	authorizationMocks "medicine/mocks/internal_/layers/business-logic/authorization"
	tags_space "medicine/mocks/internal_/layers/business-logic/simple-actions/tags-space"
	entityID "medicine/pkg/entity-id"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceListAllAvailableSA(t *testing.T) {
	t.Parallel()

	user := generators.TestUser()
	expectedTagsSpaces := []tagsSpaceModels.TagsSpace{
		generators.GenerateTagsSpace(),
		generators.GenerateTagsSpace(),
		generators.GenerateTagsSpace(),
	}
	tagsSpacesIDs := []entityID.EntityID{
		expectedTagsSpaces[0].ID,
		expectedTagsSpaces[1].ID,
		expectedTagsSpaces[2].ID,
	}

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			authorizer := authorizationMocks.NewAuthorizer(t)
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

			authorizer.EXPECT().
				AvailableResources(ctx, user, authorization.TagsSpaceResource, authorization.ReadTagsSpacePermission).
				Return(tagsSpacesIDs, nil)
			tagsSpaceAtomicActions.EXPECT().
				ListByIDs(ctx, tagsSpacesIDs).
				Return(expectedTagsSpaces, nil)

			tagsSpaces, err := sa.ListAllAvailable(ctx, user)

			assert.NoError(t, err)
			assert.Equal(t, expectedTagsSpaces, tagsSpaces)
		},
	)

	t.Run(
		"list by ids fail",
		func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			authorizer := authorizationMocks.NewAuthorizer(t)
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

			authorizer.EXPECT().
				AvailableResources(ctx, user, authorization.TagsSpaceResource, authorization.ReadTagsSpacePermission).
				Return(tagsSpacesIDs, nil)
			tagsSpaceAtomicActions.EXPECT().
				ListByIDs(ctx, tagsSpacesIDs).
				Return(nil, assert.AnError)

			tagsSpaces, err := sa.ListAllAvailable(ctx, user)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpaces)
		},
	)

	t.Run(
		"authorizer returns empty list",
		func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			authorizer := authorizationMocks.NewAuthorizer(t)
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

			authorizer.EXPECT().
				AvailableResources(ctx, user, authorization.TagsSpaceResource, authorization.ReadTagsSpacePermission).
				Return([]entityID.EntityID{}, nil)

			tagsSpaces, err := sa.ListAllAvailable(ctx, user)

			assert.NoError(t, err)
			assert.Equal(t, tagsSpaces, []tagsSpaceModels.TagsSpace{})
		},
	)

	t.Run(
		"authorizer returns error",
		func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			authorizer := authorizationMocks.NewAuthorizer(t)
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

			authorizer.EXPECT().
				AvailableResources(ctx, user, authorization.TagsSpaceResource, authorization.ReadTagsSpacePermission).
				Return(nil, assert.AnError)

			tagsSpaces, err := sa.ListAllAvailable(ctx, user)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, tagsSpaces)
		},
	)
}
