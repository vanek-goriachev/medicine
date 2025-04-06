package tag

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"medicine/internal/layers/business-logic/models/tag/identifiers"
	pkgErrors "medicine/pkg/errors/db"

	tagModels "medicine/internal/layers/business-logic/models/tag"
)

func (g *GORMGateway) GetByTagsSpaceIDAndName(
	_ context.Context,
	identifier identifiers.TagsSpaceIDAndNameIdentifier,
) (tagModels.Tag, error) {
	var dbTag Tag

	result := g.db.First(&dbTag, "tags_space_id = ? and name = ?", identifier.TagsSpaceID, identifier.Name)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tagModels.Tag{}, pkgErrors.NewDoesNotExistError(identifier)
	} else if result.Error != nil {
		return tagModels.Tag{}, fmt.Errorf("error filtering tags by tags_space_id and name: %w", result.Error)
	}

	return g.mapper.FromGORM(dbTag), nil
}
