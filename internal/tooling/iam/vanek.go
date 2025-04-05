package iam

import (
	"github.com/google/uuid"

	entity_id "medicine/pkg/entity-id"
)

// Since app has a single user and there is no real IAM - this ID will be user.
var VanekID = entity_id.EntityID(uuid.MustParse("00000000-0000-0000-0000-000000000001"))
