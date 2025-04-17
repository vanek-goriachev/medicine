package medical_file

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	medicalFileModels "medicine/internal/layers/business-logic/models/medical-file"
	gormModels "medicine/internal/layers/storage/db/gorm/models"
)

func (g *GORMGateway) Create(ctx context.Context, medicalFile medicalFileModels.MedicalFile) error {
	dbMedicalFileInfo, dbMedicalFileData := g.mapper.ToGORM(medicalFile)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error { return g.createFileInfo(ctx, dbMedicalFileInfo) })
	eg.Go(func() error { return g.createFileData(ctx, dbMedicalFileData) })

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("error on creating medicalFile: %w", err)
	}

	return nil
}

func (g *GORMGateway) createFileInfo(_ context.Context, medicalFileInfo gormModels.MedicalFileInfo) error {
	result := g.db.Create(medicalFileInfo)
	if result.Error != nil {
		return fmt.Errorf("error on creating medicalFileInfo: %w", result.Error)
	}

	return nil
}

func (g *GORMGateway) createFileData(_ context.Context, medicalFileData gormModels.MedicalFileData) error {
	result := g.db.Create(medicalFileData)
	if result.Error != nil {
		return fmt.Errorf("error on creating medicalFileData: %w", result.Error)
	}

	return nil
}
