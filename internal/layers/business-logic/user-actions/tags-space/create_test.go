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

func TestCreateTagsSpaceUA(t *testing.T) {
	t.Parallel()

	// Test data
	user := tests.TestUser()
	in := tagsSpaceUA.TagsSpaceCreateIn{
		Name: "test",
	}
	createdTagsSpace := tagsSpaceModels.TagsSpace{
		Name:   in.Name,
		Tags:   []tagModels.Tag{},
		ID:     tests.GenerateEntityID(),
		UserID: user.ID,
	}
	expectedOut := tagsSpaceUA.TagsSpaceCreateOut{
		TagsSpace: createdTagsSpace,
	}

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewCreateUA(authorizer, simpleActions)

			authorizer.EXPECT().Authorize(
				t.Context(),
				user,
				authorization.CreateTagsSpacePermission,
				authorization.TagsSpaceResource,
				"",
			).Return(nil)
			simpleActions.EXPECT().Create(
				t.Context(),
				user,
				in.Name,
			).Return(createdTagsSpace, nil)

			out, err := ua.Act(t.Context(), user, in)

			assert.NoError(t, err)
			assert.Equal(t, expectedOut, out)
		},
	)

	t.Run(
		"simple action create failed",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewCreateUA(authorizer, simpleActions)

			authorizer.EXPECT().Authorize(
				t.Context(),
				user,
				authorization.CreateTagsSpacePermission,
				authorization.TagsSpaceResource,
				"",
			).Return(nil)
			simpleActions.EXPECT().Create(
				t.Context(),
				user,
				in.Name,
			).Return(tagsSpaceModels.TagsSpace{}, assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)

	t.Run(
		"authorization failed",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewCreateUA(authorizer, simpleActions)

			authorizer.EXPECT().Authorize(
				t.Context(),
				user,
				authorization.CreateTagsSpacePermission,
				authorization.TagsSpaceResource,
				"",
			).Return(assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)
}
