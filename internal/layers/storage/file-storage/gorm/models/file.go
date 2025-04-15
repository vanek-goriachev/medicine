package models

import "github.com/google/uuid"

var FileModel = &File{} //nolint:exhaustruct // Used for ORM

type FileDataType *[]byte

type File struct {
	ID   uuid.UUID    `gorm:"primary_key;type:uuid"`
	Data FileDataType `gorm:"type:bytea"`
}
