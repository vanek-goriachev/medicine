package tags_space

import (
	"context"
	entityID "medicine/pkg/entity-id"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	userModels "medicine/pkg/user"
)

type SimpleActions interface {
	GetByID(ctx context.Context, id entityID.EntityID) (tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, user userModels.User, name string) (tagsSpaceModels.TagsSpace, error)
}
