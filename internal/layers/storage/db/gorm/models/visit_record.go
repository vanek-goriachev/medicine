package models

import (
	"time"

	"github.com/google/uuid"
)

var VisitRecordModel = &VisitRecord{}                       //nolint:exhaustruct // Used for ORM
var VisitRecordTagModel = &VisitRecordTag{}                 //nolint:exhaustruct // Used for ORM
var VisitRecordMedicalFileModel = &VisitRecordMedicalFile{} //nolint:exhaustruct // Used for ORM

type VisitRecord struct {
	Datetime time.Time `gorm:"type:timestamp"`
	Name     string    `gorm:"type:varchar(255)"`
	ID       uuid.UUID `gorm:"primary_key;type:uuid"`
}

type VisitRecordTag struct {
	VisitRecordID uuid.UUID `gorm:"type:uuid"`
	TagID         uuid.UUID `gorm:"type:uuid"`
}

type VisitRecordMedicalFile struct {
	VisitRecordID uuid.UUID `gorm:"type:uuid"`
	MedicalFileID uuid.UUID `gorm:"type:uuid"`
}
