package file

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	fileModels "medicine/internal/layers/business-logic/models/file"
	gormModels "medicine/internal/layers/storage/file-storage/gorm/models"
	entityID "medicine/pkg/entity-id"
	pkgErrors "medicine/pkg/errors/db"
)

func (g *GORMGateway) GetByID(
	_ context.Context,
	id entityID.EntityID,
) (fileModels.File, error) {
	var file gormModels.File

	result := g.db.Model(gormModels.FileModel).First(&file, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fileModels.File{}, pkgErrors.NewDoesNotExistError(id)
	} else if result.Error != nil {
		return fileModels.File{}, fmt.Errorf("error getting file by id: %w", result.Error)
	}

	return g.mapper.FromGORM(file), nil
}
