package tags_space

import (
	"context"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	customIdentifiers "medicine/internal/layers/business-logic/models/tags-space/identifiers"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type TagsSpaceAtomicActions interface {
	GetByID(ctx context.Context, id entityID.EntityID) (tagsSpaceModels.TagsSpace, error)
	ListByUserID(ctx context.Context, userID entityID.EntityID) ([]tagsSpaceModels.TagsSpace, error)
	GetByUserIDAndName(
		ctx context.Context,
		identifier customIdentifiers.UserIDAndNameIdentifier,
	) (tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, tagsSpace tagsSpaceModels.TagsSpace) error
	DeleteByID(ctx context.Context, tagsSpaceID entityID.EntityID) error
}

type TagAtomicActions interface {
	FilterByTagsSpaceID(_ context.Context, tagsSpaceID entityID.EntityID) ([]tagModels.Tag, error)
}

type TagsSpaceFactory interface {
	New(
		id entityID.EntityID,
		userID entityID.EntityID,
		name string,
		tags []tagModels.Tag,
	) (tagsSpaceModels.TagsSpace, error)
}
