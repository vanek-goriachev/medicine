package collections

import (
	gormMedicalFileMapper "medicine/internal/layers/adapters/storage/db/gorm/medical-file"
	gormTagMapper "medicine/internal/layers/adapters/storage/db/gorm/tag"
	gormTagsSpaceMapper "medicine/internal/layers/adapters/storage/db/gorm/tags-space"
	gormVisitRecordMapper "medicine/internal/layers/adapters/storage/db/gorm/visit-record"
)

type DBMappers struct {
	tag       *gormTagMapper.GORMMapper
	tagsSpace *gormTagsSpaceMapper.GORMMapper

	medicalFile *gormMedicalFileMapper.GORMMapper

	visitRecord *gormVisitRecordMapper.GORMMapper
}

func NewDBMappers() *DBMappers {
	var c DBMappers

	c.tag = gormTagMapper.NewGORMMapper()
	c.tagsSpace = gormTagsSpaceMapper.NewGORMMapper(c.tag)

	c.medicalFile = gormMedicalFileMapper.NewGORMMapper()

	c.visitRecord = gormVisitRecordMapper.NewGORMMapper()

	return &c
}
