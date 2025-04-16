package models

import (
	"github.com/google/uuid"
)

var TagModel = &Tag{} //nolint:exhaustruct // Used for ORM

type Tag struct {
	Name        string    `gorm:"type:varchar(255)"`
	ID          uuid.UUID `gorm:"primary_key;type:uuid"`
	TagsSpaceID uuid.UUID `gorm:"type:uuid;not null"` // TODO: maybe it have to be many-to-many
}
