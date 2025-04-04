package custom_identifiers

import (
	"fmt"

	entityID "medicine/pkg/entity-id"
)

type UserIDAndNameIdentifier struct {
	Name   string
	UserID entityID.EntityID
}

func (UserIDAndNameIdentifier) Identifier() {}

func (u UserIDAndNameIdentifier) String() string {
	return fmt.Sprintf(
		"identifier: user_id=%q name=%q",
		u.Name,
		u.UserID,
	)
}

func (u UserIDAndNameIdentifier) Equals(otherRaw any) bool {
	other, ok := otherRaw.(UserIDAndNameIdentifier)
	if !ok {
		return false
	}

	return u.Name == other.Name && u.UserID == other.UserID
}
