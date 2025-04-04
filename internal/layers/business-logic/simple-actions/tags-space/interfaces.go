package tags_space

import (
	"context"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	customIdentifiers "medicine/internal/tooling/identifiers/custom-identifiers"
	entityID "medicine/pkg/entity-id"
)

type AtomicActions interface {
	GetByUserIDAndName(
		ctx context.Context,
		identifier customIdentifiers.UserIDAndNameIdentifier,
	) (tagsSpaceModels.TagsSpace, error)
	Create(ctx context.Context, tagsSpace tagsSpaceModels.TagsSpace) error
}

type TagsSpaceFactory interface {
	New(id entityID.EntityID, userID entityID.EntityID, name string, tags []tagModels.Tag) tagsSpaceModels.TagsSpace
}
