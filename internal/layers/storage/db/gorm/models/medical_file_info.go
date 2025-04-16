package models

import (
	"github.com/google/uuid"
)

var MedicalFileInfoModel = &MedicalFileInfo{} //nolint:exhaustruct // Used for ORM

type MedicalFileInfo struct {
	Name string    `gorm:"type:varchar(255)"`
	ID   uuid.UUID `gorm:"primary_key;type:uuid"`
}
