package tags_space

import (
	"context"
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type TagsSpaceListAllAvailableIn struct{}

type TagsSpaceListAllAvailableOut struct {
	TagsSpaces []tagsSpaceModels.TagsSpace
}

type ListAllAvailableUA struct {
	simpleActions SimpleActions
}

func NewListAllAvailableUA(
	simpleActions SimpleActions,
) *ListAllAvailableUA {
	return &ListAllAvailableUA{
		simpleActions: simpleActions,
	}
}

func (ua *ListAllAvailableUA) Act(
	ctx context.Context,
	user userModels.User,
	_ TagsSpaceListAllAvailableIn,
) (TagsSpaceListAllAvailableOut, error) {
	tagsSpaces, err := ua.simpleActions.ListAllAvailable(ctx, user)
	if err != nil {
		return TagsSpaceListAllAvailableOut{}, fmt.Errorf("can't list all available tags spaces (ua): %w", err)
	}

	return TagsSpaceListAllAvailableOut{
		TagsSpaces: tagsSpaces,
	}, nil
}
