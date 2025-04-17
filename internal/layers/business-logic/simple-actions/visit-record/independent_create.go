package visit_record

import (
	"context"
	"fmt"
	"time"

	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
)

func (sa *SimpleActions) IndependentCreate(
	ctx context.Context,
	name string,
	datetime time.Time,
) (visitRecordModels.VisitRecord, error) {
	visitRecord, err := sa.buildVisitRecord(name, datetime)
	if err != nil {
		return visitRecordModels.VisitRecord{}, err
	}

	err = sa.atomicActions.Create(ctx, visitRecord)
	if err != nil {
		return visitRecordModels.VisitRecord{}, fmt.Errorf("can't create visitRecord: %w", err)
	}

	return visitRecord, nil
}

func (sa *SimpleActions) buildVisitRecord(name string, datetime time.Time) (visitRecordModels.VisitRecord, error) {
	id, err := sa.idGenerator.Generate()
	if err != nil {
		return visitRecordModels.VisitRecord{}, fmt.Errorf("can't generate an id: %w", err)
	}

	visitRecord, err := sa.visitRecordFactory.New(id, name, datetime)
	if err != nil {
		return visitRecordModels.VisitRecord{}, fmt.Errorf("can't build visitRecord: %w", err)
	}

	return visitRecord, nil
}
