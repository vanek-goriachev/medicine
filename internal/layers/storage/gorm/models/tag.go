package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	Name        string    `gorm:"type:varchar(255)"`
	ID          uuid.UUID `gorm:"primary_key;type:uuid"`
	TagsSpaceID uuid.UUID `gorm:"type:uuid;not null"`
}

// TODO fix linters on tests
// TODO check if i can import medicine/internal/tooling/test
