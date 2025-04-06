package entity_id

import (
	"fmt"

	"github.com/google/uuid"
)

type Generator interface {
	Generate() (EntityID, error)
}

type GeneratorImpl struct{}

func NewGenerator() *GeneratorImpl {
	return &GeneratorImpl{}
}

func (*GeneratorImpl) Generate() (EntityID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return EntityID{}, fmt.Errorf("can't generate new EntityID: %w", err)
	}

	return EntityID(id), nil
}
