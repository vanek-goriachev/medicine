package user

import (
	"context"

	entityID "medicine/pkg/entity-id"
)

func FromContext(_ context.Context) User {
	return User{
		ID: entityID.EntityID(VanekID),
	}
}
