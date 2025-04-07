package entity_id

import (
	"github.com/google/uuid"
)

type EntityID uuid.UUID

func (EntityID) Identifier() {}

func (e EntityID) Equals(otherRaw any) bool {
	other, ok := otherRaw.(EntityID)
	if !ok {
		return false
	}

	return other == e
}

func (e EntityID) String() string {
	return uuid.UUID(e).String()
}
