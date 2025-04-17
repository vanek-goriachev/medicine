package visit_record

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
	entityID "medicine/pkg/entity-id"
	pkgErrors "medicine/pkg/errors/db"
)

func (g *GORMGateway) GetByID(
	_ context.Context,
	visitRecordID entityID.EntityID,
) (visitRecordModels.VisitRecord, error) {
	var dbVisitRecord gormModels.VisitRecord

	result := g.db.Model(gormModels.VisitRecordModel).First(&dbVisitRecord, "id = ?", visitRecordID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return visitRecordModels.VisitRecord{}, pkgErrors.NewDoesNotExistError(visitRecordID)
	} else if result.Error != nil {
		return visitRecordModels.VisitRecord{}, fmt.Errorf("error getting visitRecord by id: %w", result.Error)
	}

	return g.mapper.FromGORM(dbVisitRecord), nil
}
