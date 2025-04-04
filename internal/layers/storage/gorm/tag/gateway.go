package tag

import "gorm.io/gorm"

type GORMGateway struct {
	db     *gorm.DB
	mapper tagGORMMapper
}

func NewGORMGateway(
	db *gorm.DB,
	mapper tagGORMMapper,
) *GORMGateway {
	return &GORMGateway{
		db:     db,
		mapper: mapper,
	}
}
