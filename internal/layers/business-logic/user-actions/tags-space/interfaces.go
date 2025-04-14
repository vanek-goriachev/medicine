package tags_space

import (
	"context"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

type SimpleActions interface {
	GetByID(ctx context.Context, id entityID.EntityID) (tagsSpaceModels.TagsSpace, error)
	ListByUser(ctx context.Context, user userModels.User) ([]tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, user userModels.User, name string) (tagsSpaceModels.TagsSpace, error)
	Delete(ctx context.Context, tagsSpaceID entityID.EntityID) error
}
