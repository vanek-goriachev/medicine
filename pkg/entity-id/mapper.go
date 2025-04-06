package entity_id

import (
	"fmt"

	"github.com/google/uuid"
)

type Mapper interface {
	FromString(rawID string) (EntityID, error)
}

type MapperImpl struct{}

func NewMapper() *MapperImpl {
	return &MapperImpl{}
}

func (*MapperImpl) FromString(rawID string) (EntityID, error) {
	id, err := uuid.Parse(rawID)
	if err != nil {
		return EntityID{}, fmt.Errorf("cant parse entityID from string: %w", NewInvalidEntityIDError(rawID))
	}

	return EntityID(id), nil
}
