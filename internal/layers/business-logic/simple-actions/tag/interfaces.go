package tag

import (
	"context"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	customIdentifiers "medicine/internal/layers/business-logic/models/tag/identifiers"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type AtomicActions interface {
	GetByID(ctx context.Context, tagID entityID.EntityID) (tagModels.Tag, error)
	GetByTagsSpaceIDAndName(
		ctx context.Context,
		identifier customIdentifiers.TagsSpaceIDAndNameIdentifier,
	) (tagModels.Tag, error)
	Create(ctx context.Context, tag tagModels.Tag) error
	DeleteByID(ctx context.Context, tagID entityID.EntityID) error
}

type TagFactory interface {
	New(
		id entityID.EntityID,
		tagsSpaceID entityID.EntityID,
		name string,
	) (tagModels.Tag, error)
}
