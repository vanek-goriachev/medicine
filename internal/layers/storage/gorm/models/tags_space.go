package models

import (
	"github.com/google/uuid"
)

type TagsSpace struct {
	Name   string    `gorm:"type:varchar(255);not null"`
	Tags   []Tag     `gorm:"foreignKey:TagsSpaceID;references:ID"`
	ID     uuid.UUID `gorm:"primary_key;type:uuid"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
}
