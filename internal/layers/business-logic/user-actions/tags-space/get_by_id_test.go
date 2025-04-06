package tags_space_test

import (
	"medicine/internal/layers/business-logic/authorization"
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	"medicine/internal/tooling/tests"
	authorizationMocks "medicine/mocks/internal_/layers/business-logic/authorization"
	tagsSpaceSAMock "medicine/mocks/internal_/layers/business-logic/user-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceGetByIDUA(t *testing.T) {
	t.Parallel()

	// Test data
	tagsSpaceID := tests.GenerateEntityID()
	user := tests.TestUser()
	in := tagsSpaceUA.TagsSpaceGetByIDIn{
		ID: tagsSpaceID,
	}
	expectedTagsSpace := tagsSpaceModels.TagsSpace{
		Name: "some name",
		Tags: []tagModels.Tag{
			{
				Name:        "tag1",
				ID:          tests.GenerateEntityID(),
				TagsSpaceID: tagsSpaceID,
			},
			{
				Name:        "tag2",
				ID:          tests.GenerateEntityID(),
				TagsSpaceID: tagsSpaceID,
			},
		},
		ID:     tagsSpaceID,
		UserID: user.ID,
	}
	expectedOut := tagsSpaceUA.TagsSpaceGetByIDOut{
		TagsSpace: expectedTagsSpace,
	}

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
			authorizer.EXPECT().Authorize(
				t.Context(),
				user,
				authorization.GetTagsSpacePermission,
				authorization.TagsSpaceResource,
				tagsSpaceID.String(),
			).Return(nil)

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
			authorizer.EXPECT().Authorize(
				t.Context(),
				user,
				authorization.GetTagsSpacePermission,
				authorization.TagsSpaceResource,
				tagsSpaceID.String(),
			).Return(assert.AnError)

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
