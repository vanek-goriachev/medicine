package tags_space_test

import (
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	tagsSpaceUA "medicine/internal/layers/business-logic/user-actions/tags-space"
	"medicine/internal/tooling/tests/generators"
	tagsSpaceSAMock "medicine/mocks/internal_/layers/business-logic/user-actions/tags-space"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsSpaceListAllAvailableUA(t *testing.T) {
	t.Parallel()

	// Test data
	user := generators.TestUser()

	in := tagsSpaceUA.TagsSpaceListAllAvailableIn{}
	expectedTagsSpaces := []tagsSpaceModels.TagsSpace{
		generators.GenerateTagsSpace(),
		generators.GenerateTagsSpace(),
		generators.GenerateTagsSpace(),
	}
	expectedOut := tagsSpaceUA.TagsSpaceListAllAvailableOut{
		TagsSpaces: expectedTagsSpaces,
	}

	t.Run(
		"greenpath",
		func(t *testing.T) {
			t.Parallel()

			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewListAllAvailableUA(simpleActions)

			simpleActions.EXPECT().ListAllAvailable(
				t.Context(),
				user,
			).Return(expectedTagsSpaces, nil)

			out, err := ua.Act(t.Context(), user, in)

			assert.NoError(t, err)
			assert.Equal(t, expectedOut, out)
		},
	)

	t.Run(
		"list by user fail",
		func(t *testing.T) {
			t.Parallel()

			simpleActions := tagsSpaceSAMock.NewSimpleActions(t)
			ua := tagsSpaceUA.NewListAllAvailableUA(simpleActions)

			simpleActions.EXPECT().ListAllAvailable(
				t.Context(),
				user,
			).Return([]tagsSpaceModels.TagsSpace{}, assert.AnError)

			out, err := ua.Act(t.Context(), user, in)

			assert.ErrorIs(t, err, assert.AnError)
			assert.Zero(t, out)
		},
	)

}
