package tag

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	tagModels "medicine/internal/layers/business-logic/models/tag"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
	pkgErrors "medicine/pkg/errors/db"
)

func (g *GORMGateway) GetByID(
	_ context.Context,
	id entityID.EntityID,
) (tagModels.Tag, error) {
	var tag gormModels.Tag

	result := g.db.Model(gormModels.TagModel).First(&tag, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tagModels.Tag{}, pkgErrors.NewDoesNotExistError(id)
	} else if result.Error != nil {
		return tagModels.Tag{}, fmt.Errorf("error getting tag by id: %w", result.Error)
	}

	return g.mapper.FromGORM(tag), nil
}
