package collections

import entityID "medicine/pkg/entity-id"

type CommonMappers struct {
	EntityIDMapper *entityID.MapperImpl
}

func NewCommonMappers() *CommonMappers {
	var c CommonMappers

	c.EntityIDMapper = entityID.NewMapper()

	return &c
}
