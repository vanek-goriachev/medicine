package visit_record

import "gorm.io/gorm"

type GORMGateway struct {
	db     *gorm.DB
	mapper visitRecordGORMMapper
}

func NewGORMGateway(
	db *gorm.DB,
	mapper visitRecordGORMMapper,
) *GORMGateway {
	return &GORMGateway{
		db:     db,
		mapper: mapper,
	}
}
