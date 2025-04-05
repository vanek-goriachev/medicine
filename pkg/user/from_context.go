package user

import (
	"context"

	entityID "medicine/pkg/entity-id"
)

func FromContext(_ context.Context) (User, error) { //nolint:unparam // Gonna fix later
	return User{
		ID:          entityID.EntityID(VanekID),
		isAnonymous: false,
	}, nil
}
