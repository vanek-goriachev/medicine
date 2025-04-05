package user

import (
	entityID "medicine/pkg/entity-id"
)

type User struct {
	ID entityID.EntityID

	IsAnonymous bool
}
