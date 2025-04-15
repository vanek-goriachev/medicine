package tags_space

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
	pkgErrors "medicine/pkg/errors/db"
)

func (g *GORMGateway) GetByID(
	_ context.Context,
	id entityID.EntityID,
) (tagsSpaceModels.TagsSpace, error) {
	var tagsSpace gormModels.TagsSpace

	result := g.db.Model(gormModels.TagsSpaceModel).Preload(gormModels.TagsField).First(&tagsSpace, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tagsSpaceModels.TagsSpace{}, pkgErrors.NewDoesNotExistError(id)
	} else if result.Error != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("error getting tagsSpace by id: %w", result.Error)
	}

	return g.mapper.FromGORM(tagsSpace), nil
}
