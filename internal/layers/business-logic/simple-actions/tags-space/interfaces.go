package tags_space

import (
	"context"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type AtomicActions interface {
	GetByID(ctx context.Context, id entityID.EntityID) (tagsSpaceModels.TagsSpace, error)
	ListByUserID(ctx context.Context, userID entityID.EntityID) ([]tagsSpaceModels.TagsSpace, error)
	GetByUserIDAndName(
		ctx context.Context,
		identifier customIdentifiers.UserIDAndNameIdentifier,
	) (tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, tagsSpace tagsSpaceModels.TagsSpace) error
}

type TagsSpaceFactory interface {
	New(
		id entityID.EntityID,
		userID entityID.EntityID,
		name string,
		tags []tagModels.Tag,
	) (tagsSpaceModels.TagsSpace, error)
}
