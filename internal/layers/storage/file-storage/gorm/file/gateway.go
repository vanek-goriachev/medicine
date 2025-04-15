package file

import "gorm.io/gorm"

type GORMGateway struct {
	db     *gorm.DB
	mapper fileGORMMapper
}

func NewGORMGateway(
	db *gorm.DB,
	mapper fileGORMMapper,
) *GORMGateway {
	return &GORMGateway{
		db:     db,
		mapper: mapper,
	}
}
