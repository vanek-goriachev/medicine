package tags_space

import (
	"context"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type SimpleActions interface {
	Create(ctx context.Context, user userModels.User, name string) (tagsSpaceModels.TagsSpace, error)
}
