package visit_record

import (
	"context"
	"fmt"

	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
)

func (g *GORMGateway) Create(_ context.Context, visitRecord visitRecordModels.VisitRecord) error {
	dbVisitRecord := g.mapper.ToGORM(visitRecord)

	result := g.db.Create(&dbVisitRecord)
	if result.Error != nil {
		return fmt.Errorf("error on creating visitRecord: %w", result.Error)
	}

	return nil
}
