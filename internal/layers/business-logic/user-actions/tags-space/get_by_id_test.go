package tags_space_test

import (
	"medicine/internal/layers/business-logic/authorization"
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	authorizationMocks "medicine/mocks/internal_/layers/business-logic/authorization"
	tagsSpaceSAMock "medicine/mocks/internal_/layers/business-logic/user-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceGetByIDUA(t *testing.T) {
	t.Parallel()

	// Test data
	tagsSpaceID := generators.GenerateEntityID()
	user := generators.TestUser()
	in := &tagsSpaceUA.TagsSpaceGetByIDIn{
		ID: tagsSpaceID,
	}
	expectedTagsSpace := tagsSpaceModels.TagsSpace{
		Name: "some name",
		Tags: []tagModels.Tag{
			{
				Name:        "tag1",
				ID:          generators.GenerateEntityID(),
				TagsSpaceID: tagsSpaceID,
			},
			{
				Name:        "tag2",
				ID:          generators.GenerateEntityID(),
				TagsSpaceID: tagsSpaceID,
			},
		},
		ID: tagsSpaceID,
	}
	expectedOut := tagsSpaceUA.TagsSpaceGetByIDOut{
		TagsSpace: expectedTagsSpace,
	}
	authAction := authorization.NewAction(
		authorization.ReadTagsSpacePermission,
		authorization.TagsSpaceResource,
		tagsSpaceID,
	)

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewGetByIDUA(authorizer, simpleActions)

			simpleActions.EXPECT().GetByID(
				t.Context(),
				in.ID,
			).Return(expectedTagsSpace, nil)
			authorizer.EXPECT().Authorize(t.Context(), user, authAction).Return(nil)

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
			ua := tagsSpaceUA.NewGetByIDUA(authorizer, simpleActions)

			simpleActions.EXPECT().GetByID(
				t.Context(),
				in.ID,
			).Return(expectedTagsSpace, nil)
			authorizer.EXPECT().Authorize(t.Context(), user, authAction).Return(assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)

	t.Run(
		"get fail",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewGetByIDUA(authorizer, simpleActions)

			simpleActions.EXPECT().GetByID(
				t.Context(),
				in.ID,
			).Return(tagsSpaceModels.TagsSpace{}, assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)

}
