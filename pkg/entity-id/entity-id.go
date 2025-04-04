package entity_id

import (
	"github.com/google/uuid"
)

type EntityID uuid.UUID

func (e EntityID) String() string {
	return uuid.UUID(e).String()
}
