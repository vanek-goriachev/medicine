package collections

import (
	gormTagMapper "medicine/internal/layers/adapters/storage/db/gorm/tag"
	gormTagsSpaceMapper "medicine/internal/layers/adapters/storage/db/gorm/tags-space"
)

type DBMappers struct {
	tag       *gormTagMapper.GORMMapper
	tagsSpace *gormTagsSpaceMapper.GORMMapper
}

func NewDBMappers() *DBMappers {
	var c DBMappers

	c.tag = gormTagMapper.NewGORMMapper()
	c.tagsSpace = gormTagsSpaceMapper.NewGORMMapper(c.tag)

	return &c
}
