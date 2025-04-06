package tags_space

import (
	"context"
	"errors"
	"fmt"
	customIdentifiers "medicine/internal/layers/business-logic/models/tags-space/identifiers"

	"gorm.io/gorm"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
	pkgErrors "medicine/pkg/errors/db"
)

func (g *GORMGateway) GetByUserIDAndName(
	_ context.Context,
	identifier customIdentifiers.UserIDAndNameIdentifier,
) (tagsSpaceModels.TagsSpace, error) {
	var tagsSpace TagsSpace

	result := g.db.First(&tagsSpace, "user_id = ? and name = ?", identifier.UserID, identifier.Name)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tagsSpaceModels.TagsSpace{}, pkgErrors.NewDoesNotExistError(identifier)
	} else if result.Error != nil {
		return tagsSpaceModels.TagsSpace{}, fmt.Errorf("error getting tagsSpace by user_id and name: %w", result.Error)
	}

	return g.mapper.FromGORM(tagsSpace), nil
}
