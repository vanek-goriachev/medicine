package user

import (
	"github.com/google/uuid"

	entityID "medicine/pkg/entity-id"
)

type User struct {
	ID entityID.EntityID
}

//nolint:gochecknoglobals // Since app has a single user and there is no real IAM - this ID will be user
var VanekID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
