package collections

import (
	"medicine/internal/tooling/datetime"
	entityID "medicine/pkg/entity-id"
	"time"
)

type CommonMappers struct {
	EntityIDMapper entityID.Mapper
	Datetime       datetime.Mapper
}

func NewCommonMappers() *CommonMappers {
	var c CommonMappers

	c.EntityIDMapper = entityID.NewMapper()
	c.Datetime = datetime.NewMapper(time.RFC3339)

	return &c
}
