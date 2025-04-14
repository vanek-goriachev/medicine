package tag

import (
	"context"
	entityID "medicine/pkg/entity-id"
)

func (sa *SimpleActions) UntagAll(
	ctx context.Context,
	id entityID.EntityID,
) error {
	return nil
}
