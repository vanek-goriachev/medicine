package tags_space

import (
	"github.com/google/uuid"

	tagsDB "medicine/internal/layers/storage/gorm/tag"
)

type TagsSpace struct {
	Name   string       `gorm:"type:varchar(255);not null"`
	Tags   []tagsDB.Tag `gorm:"foreignKey:TagsSpaceID;references:ID"`
	ID     uuid.UUID    `gorm:"primary_key;type:uuid"`
	UserID uuid.UUID    `gorm:"type:uuid;not null"`
}
