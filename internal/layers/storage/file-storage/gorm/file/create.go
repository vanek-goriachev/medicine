package file

import (
	"context"
	"fmt"

	fileModels "medicine/internal/layers/business-logic/models/file"
)

func (g *GORMGateway) Create(_ context.Context, file fileModels.File) error {
	dbFile := g.mapper.ToGORM(file)

	result := g.db.Create(&dbFile)

	if result.Error != nil {
		return fmt.Errorf("error on creating file: %w", result.Error)
	}

	return nil
}
