package medical_file

import "gorm.io/gorm"

type GORMGateway struct {
	db     *gorm.DB
	mapper medicalFileGORMMapper
}

func NewGORMGateway(
	db *gorm.DB,
	mapper medicalFileGORMMapper,
) *GORMGateway {
	return &GORMGateway{
		db:     db,
		mapper: mapper,
	}
}
