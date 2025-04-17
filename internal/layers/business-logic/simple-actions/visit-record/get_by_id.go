package visit_record

import (
	"context"
	"fmt"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
)

func (sa *SimpleActions) GetByID(
	ctx context.Context,
	visitRecordID entityID.EntityID,
) (visitRecordModels.VisitRecord, error) {
	visitRecord, err := sa.atomicActions.GetByID(ctx, visitRecordID)
	if err != nil {
		return visitRecordModels.VisitRecord{}, fmt.Errorf("can't get visit record by id (sa): %w", err)
	}

	return visitRecord, nil
}
