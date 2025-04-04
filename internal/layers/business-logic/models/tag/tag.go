package tag

import (
	entityID "medicine/pkg/entity-id"
)

type Tag struct {
	Name        string
	ID          entityID.EntityID
	TagsSpaceID entityID.EntityID
}
