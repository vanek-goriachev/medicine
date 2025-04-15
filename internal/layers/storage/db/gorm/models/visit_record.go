package models

import (
	"github.com/google/uuid"
)

var VisitRecordModel = &VisitRecord{} //nolint:exhaustruct // Used for ORM

type VisitRecord struct {
	Name   string    `gorm:"type:varchar(255)"`
	ID     uuid.UUID `gorm:"primary_key;type:uuid"`
	FileID uuid.UUID `gorm:"type:uuid"`
}

type VisitRecordTags struct {
}
