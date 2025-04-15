package collections

import (
	"medicine/internal/appcore/dependencies/db"

	tagGORM "medicine/internal/layers/storage/db/gorm/tag"
	tagsSpaceGORM "medicine/internal/layers/storage/db/gorm/tags-space"
)

type DBGateways struct {
	tag       *tagGORM.GORMGateway
	tagsSpace *tagsSpaceGORM.GORMGateway
}

func NewDBGateways(db *db.DB, gormMappers *DBMappers) *DBGateways {
	var c DBGateways

	c.tag = tagGORM.NewGORMGateway(db.GormDB, gormMappers.tag)
	c.tagsSpace = tagsSpaceGORM.NewGORMGateway(db.GormDB, gormMappers.tagsSpace)

	return &c
}
