package visit_record

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	visitRecordModels "medicine/internal/layers/business-logic/models/visit-record"
	entityID "medicine/pkg/entity-id"
	"time"
)

func (sa *SimpleActions) CreateWithEntities(
	ctx context.Context,
	name string,
	datetime time.Time,

	uploadedMedicalFiles []medicalFileModels.UploadedMedicalFile,
	tagIDs []entityID.EntityID,
) (visitRecordModels.VisitRecord, visitRecordModels.VisitRecordLinkedEntities, error) {
	visitRecord, files, err := sa.createNewEntities(ctx, name, datetime, uploadedMedicalFiles)
	if err != nil {
		return visitRecordModels.VisitRecord{}, visitRecordModels.VisitRecordLinkedEntities{}, err
	}

	linkedEntities, err := sa.linkEntities(
		ctx,
		visitRecord,
		files,
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
	uploadedMedicalFiles []medicalFileModels.UploadedMedicalFile,
) (visitRecordModels.VisitRecord, []medicalFileModels.MedicalFile, error) {
	var visitRecord visitRecordModels.VisitRecord
	var files []medicalFileModels.MedicalFile

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

	eg.Go(
		func() error {
			var err error
			files, err = sa.fileAtomicActions.CreateMultiple(ctx, uploadedMedicalFiles)
			if err != nil {
				return fmt.Errorf("failed to create medical files: %w", err)
			}
			return nil
		},
	)

	if err := eg.Wait(); err != nil {
		return visitRecordModels.VisitRecord{}, nil, fmt.Errorf("failed to create new entities: %w", err)
	}

	return visitRecord, files, nil
}

func (sa *SimpleActions) linkEntities(
	ctx context.Context,
	visitRecord visitRecordModels.VisitRecord,
	medicalFiles []medicalFileModels.MedicalFile,
	tagIDs []entityID.EntityID,
) (visitRecordModels.VisitRecordLinkedEntities, error) {
	eg, ctx := errgroup.WithContext(ctx)

	medicalFileIDs := sa.extractMedicalFileIDs(medicalFiles)

	eg.Go(func() error { return sa.linkMedicalFiles(ctx, visitRecord, medicalFileIDs) })

	eg.Go(func() error { return sa.linkTags(ctx, visitRecord, tagIDs) })

	if err := eg.Wait(); err != nil {
		return visitRecordModels.VisitRecordLinkedEntities{}, fmt.Errorf("failed to link entities: %w", err)
	}

	return visitRecordModels.VisitRecordLinkedEntities{
		MedicalFileIDs: medicalFileIDs,
		TagIDs:         tagIDs,
	}, nil
}

func (sa *SimpleActions) extractMedicalFileIDs(medicalFiles []medicalFileModels.MedicalFile) []entityID.EntityID {
	medicalFileIDs := make([]entityID.EntityID, len(medicalFiles))
	for i, file := range medicalFiles {
		medicalFileIDs[i] = file.ID
	}
	return medicalFileIDs
}

func (sa *SimpleActions) linkMedicalFiles(
	ctx context.Context,
	visitRecord visitRecordModels.VisitRecord,
	medicalFileIDs []entityID.EntityID,
) error {
	err := sa.atomicActions.LinkMedicalFiles(ctx, visitRecord.ID, medicalFileIDs)
	if err != nil {
		return fmt.Errorf("failed to link medical files: %w", err)
	}

	return nil
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
