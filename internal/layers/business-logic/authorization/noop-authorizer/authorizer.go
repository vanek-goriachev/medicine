package noop_authorizer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"medicine/internal/layers/business-logic/authorization"
	"medicine/internal/layers/storage/db/gorm/models"
	models2 "medicine/internal/layers/storage/file-storage/gorm/models"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

// NoopAuthorizer is an authorizer which grants access to all resources and actions.
type NoopAuthorizer struct {
	gormDB          *gorm.DB
	gormFileStorage *gorm.DB
}

func NewNoopAuthorizer(
	gormDB *gorm.DB,
	gormFileStorage *gorm.DB,
) *NoopAuthorizer {
	return &NoopAuthorizer{
		gormDB:          gormDB,
		gormFileStorage: gormFileStorage,
	}
}

func (*NoopAuthorizer) Authorize(
	_ context.Context,
	_ userModels.User,
	_ authorization.Action,
) error {
	return nil
}

func (*NoopAuthorizer) BatchAuthorize(
	_ context.Context,
	_ userModels.User,
	_ []authorization.Action,
) error {
	return nil
}

func (na *NoopAuthorizer) AvailableResources(
	_ context.Context,
	_ userModels.User,
	resource authorization.Resource,
	_ authorization.Permission,
) ([]entityID.EntityID, error) {
	var model any
	var db *gorm.DB

	switch resource {
	case authorization.TagsSpaceResource:
		model = models.TagsSpaceModel
		db = na.gormDB
	case authorization.TagResource:
		model = models.TagModel
		db = na.gormDB
	case authorization.FileResource:
		model = models2.FileModel
		db = na.gormFileStorage
	default:
		return nil, NewInvalidResourceError(resource)
	}

	rawIDs, err := scanIDs(db, model)
	if err != nil {
		return nil, fmt.Errorf("error when noop-authorizing on resource='%v': %w", resource, err)
	}

	ids := convertIDs(rawIDs)

	return ids, nil
}

func scanIDs(db *gorm.DB, model any) ([]uuid.UUID, error) {
	var ids []uuid.UUID

	result := db.Model(model).Select("id").Find(&ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return ids, nil
}

func convertIDs(ids []uuid.UUID) []entityID.EntityID {
	result := make([]entityID.EntityID, len(ids))

	for i, id := range ids {
		result[i] = entityID.EntityID(id)
	}

	return result
}
