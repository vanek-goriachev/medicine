package tags_space_test

import (
	"medicine/internal/layers/business-logic/authorization"
	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	authorizationMocks "medicine/mocks/internal_/layers/business-logic/authorization"
	tagsSpaceSAMock "medicine/mocks/internal_/layers/business-logic/user-actions/tags-space"
	entityID "medicine/pkg/entity-id"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTagsSpaceUA(t *testing.T) {
	t.Parallel()

	// Test data
	user := generators.TestUser()
	in := &tagsSpaceUA.TagsSpaceCreateIn{
		Name: "test",
	}
	createdTagsSpace := tagsSpaceModels.TagsSpace{
		Name: in.Name,
		Tags: []tagModels.Tag{},
		ID:   generators.GenerateEntityID(),
	}
	expectedOut := tagsSpaceUA.TagsSpaceCreateOut{
		TagsSpace: createdTagsSpace,
	}
	authAction := authorization.NewAction(
		authorization.CreateTagsSpacePermission,
		authorization.TagsSpaceResource,
		entityID.EntityID{},
	)

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			authorizer := authorizationMocks.NewAuthorizer(t)
			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewCreateUA(authorizer, simpleActions)

			authorizer.EXPECT().Authorize(t.Context(), user, authAction).Return(nil)
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

			authorizer.EXPECT().Authorize(t.Context(), user, authAction).Return(nil)
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

			authorizer.EXPECT().Authorize(t.Context(), user, authAction).Return(assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)
}
