package noop_authorizer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"medicine/internal/layers/business-logic/authorization"
	"medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
	userModels "medicine/pkg/user"
)

// NoopAuthorizer is an authorizer which grants access to all resources and actions.
type NoopAuthorizer struct {
	gormDB *gorm.DB
}

func NewNoopAuthorizer(
	gormDB *gorm.DB,
) *NoopAuthorizer {
	return &NoopAuthorizer{
		gormDB: gormDB,
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
	mapping := map[authorization.Resource]any{
		authorization.TagsSpaceResource:   models.TagsSpaceModel,
		authorization.TagResource:         models.TagModel,
		authorization.MedicalFileResource: models.MedicalFileInfoModel,
		authorization.VisitRecordResource: models.VisitRecordModel,
	}

	model, ok := mapping[resource]
	if !ok {
		return nil, NewInvalidResourceError(resource)
	}

	rawIDs, err := na.scanIDs(model)
	if err != nil {
		return nil, fmt.Errorf("error when noop-authorizing on resource='%v': %w", resource, err)
	}

	ids := convertIDs(rawIDs)

	return ids, nil
}

func (na *NoopAuthorizer) scanIDs(model any) ([]uuid.UUID, error) {
	var ids []uuid.UUID

	result := na.gormDB.Model(model).Select("id").Find(&ids)
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
