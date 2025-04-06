package tag

import (
	"context"
	customIdentifiers "medicine/internal/layers/business-logic/models/tag/identifiers"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type AtomicActions interface {
	GetByTagsSpaceIDAndName(
		ctx context.Context,
		identifier customIdentifiers.TagsSpaceIDAndNameIdentifier,
	) (tagModels.Tag, error)
	Create(ctx context.Context, tag tagModels.Tag) error
}

type TagFactory interface {
	New(
		id entityID.EntityID,
		tagsSpaceID entityID.EntityID,
		name string,
	) (tagModels.Tag, error)
}
