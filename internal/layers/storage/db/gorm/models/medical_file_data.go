package models

import "github.com/google/uuid"

var MedicalFileDataModel = &MedicalFileData{} //nolint:exhaustruct // Used for ORM

type DataType []byte

type MedicalFileData struct {
	Data DataType  `gorm:"type:bytea"`
	ID   uuid.UUID `gorm:"primary_key;type:uuid"`
}
