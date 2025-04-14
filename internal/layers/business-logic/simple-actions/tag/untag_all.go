package tag

import (
	"context"

	entityID "medicine/pkg/entity-id"
)

func (*SimpleActions) UntagAll(
	_ context.Context,
	_ entityID.EntityID,
) error {
	return nil
}
