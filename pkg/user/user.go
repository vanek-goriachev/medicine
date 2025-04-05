package user

import (
	"github.com/google/uuid"

	entityID "medicine/pkg/entity-id"
)

// Since app has a single user and there is no real IAM - this ID will be user.
var VanekID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

type User struct {
	ID entityID.EntityID

	isAnonymous bool
}

func (u *User) IsAnonymous() bool {
	return u.isAnonymous
}
