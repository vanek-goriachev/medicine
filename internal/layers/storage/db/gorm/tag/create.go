package tag

import (
	"context"
	"fmt"

	tagModels "medicine/internal/layers/business-logic/models/tag"
)

func (g *GORMGateway) Create(_ context.Context, tag tagModels.Tag) error {
	dbTag := g.mapper.ToGORM(tag)

	result := g.db.Create(&dbTag)
	if result.Error != nil {
		return fmt.Errorf("error on creating tag: %w", result.Error)
	}

	return nil
}
