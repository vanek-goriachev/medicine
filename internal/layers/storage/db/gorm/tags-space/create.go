package tags_space

import (
	"context"
	"fmt"

	tagsSpaceModels "medicine/internal/layers/business-logic/models/tags-space"
)

func (g *GORMGateway) Create(_ context.Context, tagsSpace tagsSpaceModels.TagsSpace) error {
	dbTagsSpace := g.mapper.ToGORM(tagsSpace)

	result := g.db.Create(&dbTagsSpace)

	if result.Error != nil {
		return fmt.Errorf("error on creating tagsSpace: %w", result.Error)
	}

	return nil
}
