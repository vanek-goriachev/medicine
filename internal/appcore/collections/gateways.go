package collections

import (
	"gorm.io/gorm"

	tagGORM "medicine/internal/layers/storage/db/gorm/tag"
	tagsSpaceGORM "medicine/internal/layers/storage/db/gorm/tags-space"
)

type Gateways struct {
	tag       *tagGORM.GORMGateway
	tagsSpace *tagsSpaceGORM.GORMGateway
}

func NewGateways(db *gorm.DB, gormMappers *GORMMappers) *Gateways {
	var c Gateways

	c.tag = tagGORM.NewGORMGateway(db, gormMappers.tag)
	c.tagsSpace = tagsSpaceGORM.NewGORMGateway(db, gormMappers.tagsSpace)

	return &c
}
