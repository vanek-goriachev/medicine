package tags_space

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	customIdentifiers "medicine/internal/layers/business-logic/models/tags-space/identifiers"
	gormModels "medicine/internal/layers/storage/gorm/models"
	pkgErrors "medicine/pkg/errors/db"
)

func (g *GORMGateway) GetByUserIDAndName(
	_ context.Context,
	identifier customIdentifiers.UserIDAndNameIdentifier,
) (tagsSpaceModels.TagsSpace, error) {
	var tagsSpace gormModels.TagsSpace

	result := g.db.Model(gormModels.TagsSpaceModel).
		Preload(gormModels.TagsField).
		First(&tagsSpace, "user_id = ? and name = ?", identifier.UserID, identifier.Name)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tagsSpaceModels.TagsSpace{}, pkgErrors.NewDoesNotExistError(identifier)
	} else if result.Error != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf(
			"error getting tagsSpace by user_id and name: %w",
			result.Error,
		)
	}

	return g.mapper.FromGORM(tagsSpace), nil
}
