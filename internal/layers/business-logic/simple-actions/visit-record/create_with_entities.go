package visit_record

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
	"time"
)

func (sa *SimpleActions) CreateWithEntities(
	ctx context.Context,
	name string,
	datetime time.Time,

	tagIDs []entityID.EntityID,
) (visitRecordModels.VisitRecord, visitRecordModels.VisitRecordLinkedEntities, error) {
	visitRecord, err := sa.createNewEntities(ctx, name, datetime)
	if err != nil {
		return visitRecordModels.VisitRecord{}, visitRecordModels.VisitRecordLinkedEntities{}, err
	}

	linkedEntities, err := sa.linkEntities(
		ctx,
		visitRecord,
		tagIDs,
	)
	if err != nil {
		return visitRecordModels.VisitRecord{}, visitRecordModels.VisitRecordLinkedEntities{}, err
	}

	return visitRecord, linkedEntities, nil
}

func (sa *SimpleActions) createNewEntities(
	ctx context.Context,
	name string,
	datetime time.Time,
) (visitRecordModels.VisitRecord, error) {
	var visitRecord visitRecordModels.VisitRecord

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(
		func() error {
			var err error
			visitRecord, err = sa.IndependentCreate(ctx, name, datetime)
			if err != nil {
				return fmt.Errorf("failed to create visit record: %w", err)
			}
			return nil
		},
	)

	if err := eg.Wait(); err != nil {
		return visitRecordModels.VisitRecord{}, fmt.Errorf("failed to create new entities: %w", err)
	}

	return visitRecord, nil
}

func (sa *SimpleActions) linkEntities(
	ctx context.Context,
	visitRecord visitRecordModels.VisitRecord,
	tagIDs []entityID.EntityID,
) (visitRecordModels.VisitRecordLinkedEntities, error) {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error { return sa.linkTags(ctx, visitRecord, tagIDs) })

	if err := eg.Wait(); err != nil {
		return visitRecordModels.VisitRecordLinkedEntities{}, fmt.Errorf("failed to link entities: %w", err)
	}

	return visitRecordModels.VisitRecordLinkedEntities{
		MedicalFileIDs: []entityID.EntityID{},
		TagIDs:         tagIDs,
	}, nil
}

func (sa *SimpleActions) linkTags(
	ctx context.Context,
	visitRecord visitRecordModels.VisitRecord,
	tagIDs []entityID.EntityID,
) error {
	err := sa.atomicActions.LinkTags(ctx, visitRecord.ID, tagIDs)
	if err != nil {
		return fmt.Errorf("failed to link tags: %w", err)
	}

	return nil
}
