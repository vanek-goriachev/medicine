package collections

import (
	"medicine/internal/appcore/dependencies/db"
	medicalFileGORM "medicine/internal/layers/storage/db/gorm/medical-file"
	tagGORM "medicine/internal/layers/storage/db/gorm/tag"
	tagsSpaceGORM "medicine/internal/layers/storage/db/gorm/tags-space"
	visitRecordGORM "medicine/internal/layers/storage/db/gorm/visit-record"
)

type DBGateways struct {
	tag       *tagGORM.GORMGateway
	tagsSpace *tagsSpaceGORM.GORMGateway

	medicalFile *medicalFileGORM.GORMGateway

	visitRecord *visitRecordGORM.GORMGateway
}

func NewDBGateways(database *db.DB, gormMappers *DBMappers) *DBGateways {
	var c DBGateways

	c.tag = tagGORM.NewGORMGateway(database.GormDB, gormMappers.tag)
	c.tagsSpace = tagsSpaceGORM.NewGORMGateway(database.GormDB, gormMappers.tagsSpace)

	c.medicalFile = medicalFileGORM.NewGORMGateway(database.GormDB, gormMappers.medicalFile)

	c.visitRecord = visitRecordGORM.NewGORMGateway(database.GormDB, gormMappers.visitRecord)

	return &c
}
