package tags_space_test

import (
	"medicine/internal/layers/business-logic/authorization"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	authorizationMocks "medicine/mocks/internal_/layers/business-logic/authorization"
	tagsSpaceSAMock "medicine/mocks/internal_/layers/business-logic/user-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceListByUserUA(t *testing.T) {
	t.Parallel()

	// Test data
	user := generators.TestUser()

	in := tagsSpaceUA.TagsSpaceListByUserIn{}
	expectedTagsSpaces := []tagsSpaceModels.TagsSpace{
		generators.GenerateTagsSpace(user.ID),
		generators.GenerateTagsSpace(user.ID),
		generators.GenerateTagsSpace(user.ID),
	}
	expectedOut := tagsSpaceUA.TagsSpaceListByUserOut{
		TagsSpaces: expectedTagsSpaces,
	}
	authActions := []authorization.Action{
		authorization.NewAction(authorization.ReadTagsSpacePermission, authorization.TagsSpaceResource, expectedTagsSpaces[0].ID.String()),
		authorization.NewAction(authorization.ReadTagsSpacePermission, authorization.TagsSpaceResource, expectedTagsSpaces[1].ID.String()),
		authorization.NewAction(authorization.ReadTagsSpacePermission, authorization.TagsSpaceResource, expectedTagsSpaces[2].ID.String()),
	}

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewListByUserUA(authorizer, simpleActions)

			simpleActions.EXPECT().ListByUser(
				t.Context(),
				user,
			).Return(expectedTagsSpaces, nil)
			authorizer.EXPECT().BatchAuthorize(t.Context(), user, authActions).Return(nil)

			out, err := ua.Act(t.Context(), user, in)

			assert.NoError(t, err)
			assert.Equal(t, expectedOut, out)
		},
	)

	t.Run(
		"authorization fail",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewListByUserUA(authorizer, simpleActions)

			simpleActions.EXPECT().ListByUser(
				t.Context(),
				user,
			).Return(expectedTagsSpaces, nil)
			authorizer.EXPECT().BatchAuthorize(t.Context(), user, authActions).Return(assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)

	t.Run(
		"list by user fail",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewListByUserUA(authorizer, simpleActions)

			simpleActions.EXPECT().ListByUser(
				t.Context(),
				user,
			).Return([]tagsSpaceModels.TagsSpace{}, assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)

}
