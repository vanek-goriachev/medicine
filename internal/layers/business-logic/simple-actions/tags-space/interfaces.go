package tags_space

import (
	"context"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type TagsSpaceAtomicActions interface {
	GetByID(ctx context.Context, id entityID.EntityID) (tagsSpaceModels.TagsSpace, error)
	ListByIDs(ctx context.Context, ids []entityID.EntityID) ([]tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, tagsSpace tagsSpaceModels.TagsSpace) error
	DeleteByID(ctx context.Context, tagsSpaceID entityID.EntityID) error
}

type TagAtomicActions interface {
	FilterByTagsSpaceID(_ context.Context, tagsSpaceID entityID.EntityID) ([]tagModels.Tag, error)
}

type TagsSpaceFactory interface {
	New(
		id entityID.EntityID,
		name string,
		tags []tagModels.Tag,
	) (tagsSpaceModels.TagsSpace, error)
}
