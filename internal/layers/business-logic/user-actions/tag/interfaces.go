package tag

import (
	"context"
	tagModels "medicine/internal/layers/business-logic/models/tag"

	entityID "medicine/pkg/entity-id"
)

type SimpleActions interface {
	Create(ctx context.Context, name string, tagsSpaceID entityID.EntityID) (tagModels.Tag, error)
}
