package generators

import (
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

func TestUser() userModels.User {
	return userModels.User{
		ID:          userModels.VanekID,
		IsAnonymous: false,
	}
}

func AnonymousTestUser() userModels.User {
	return userModels.User{
		ID:          entityID.EntityID{},
		IsAnonymous: true,
	}
}
