package models

import "github.com/google/uuid"

var FileModel = &File{} //nolint:exhaustruct // Used for ORM

type File struct {
	ID   uuid.UUID `gorm:"primary_key;type:uuid"`
	Name string    `gorm:"type:varchar(255)"`
	Data FileData  `gorm:"type:bytea"`
}
