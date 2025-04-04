package tags_space

import (
	"context"
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type CreateTagsSpaceIn struct {
	Name string
}

type CreateTagsSpaceOut struct {
	TagsSpace tagsSpaceModels.TagsSpace
}

type CreateUA struct {
	simpleActions SimpleActions
}

func NewCreateUA(simpleActions SimpleActions) *CreateUA {
	return &CreateUA{
		simpleActions: simpleActions,
	}
}

func (ua *CreateUA) Act(ctx context.Context, user userModels.User, in CreateTagsSpaceIn) (CreateTagsSpaceOut, error) {
	tagsSpace, err := ua.simpleActions.Create(ctx, user, in.Name)
	if err != nil {
		return CreateTagsSpaceOut{}, fmt.Errorf("can't create tags space: %w", err)
	}

	return CreateTagsSpaceOut{
		TagsSpace: tagsSpace,
	}, nil
}
