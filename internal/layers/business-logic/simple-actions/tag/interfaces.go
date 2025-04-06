package tag

import (
	"context"
	customIdentifiers "medicine/internal/layers/business-logic/models/tag/identifiers"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	entityID "medicine/pkg/entity-id"
)

type EntityIDGenerator interface {
	Generate() (entityID.EntityID, error)
}

type AtomicActions interface {
	GetByTagsSpaceIDAndName(
		ctx context.Context,
		identifier customIdentifiers.TagsSpaceIDAndNameIdentifier,
	) (tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, tag tagModels.Tag) error
}

type TagFactory interface {
	New(
		id entityID.EntityID,
		name string,
		tagsSpaceID entityID.EntityID,
	) (tagModels.Tag, error)
}
