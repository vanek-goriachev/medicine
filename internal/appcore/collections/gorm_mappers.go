package collections

import (
	gormTagMapper "medicine/internal/layers/adapters/storage/gorm/tag"
	gormTagsSpaceMapper "medicine/internal/layers/adapters/storage/gorm/tags-space"
)

type GORMMappers struct {
	tag       *gormTagMapper.GORMMapper
	tagsSpace *gormTagsSpaceMapper.GORMMapper
}

func NewGORMMappers() *GORMMappers {
	var c GORMMappers

	c.tag = gormTagMapper.NewGORMMapper()
	c.tagsSpace = gormTagsSpaceMapper.NewGORMMapper(c.tag)

	return &c
}
