package tags_space

import (
	"gorm.io/gorm"
)

type GORMGateway struct {
	db     *gorm.DB
	mapper tagsSpaceGORMMapper
}

func NewGORMGateway(
	db *gorm.DB,
	mapper tagsSpaceGORMMapper,
) *GORMGateway {
	return &GORMGateway{
		db:     db,
		mapper: mapper,
	}
}
