package tag

import (
	"context"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type SimpleActions interface {
	Create(ctx context.Context, user userModels.User, name string) (tagsSpaceModels.TagsSpace, error)
	FindSimilarWithinTagsSpace(
		ctx context.Context,
		tagsSpaceID entityID.EntityID,
		name string,
	) ([]tagsSpaceModels.TagsSpace, error)
}
