package entity_id

import (
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
		return EntityID{}, NewCantParsedEntityIDError(rawID)
	}

	return EntityID(id), nil
}
