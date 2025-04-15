package file_storage

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormModels "medicine/internal/layers/storage/file-storage/gorm/models"
	"medicine/pkg/retry"
	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

type FileStorage struct {
	GormFileStorage *gorm.DB
	logger          *slog.Logger
}

func NewFileStorage(
	ctx context.Context,
	fileStorageCfg Config, //nolint:gocritic // Its fine to pass config by value
	logger *slog.Logger,
) (*FileStorage, error) {
	var fs FileStorage
	fs.logger = logger

	err := fs.connectFileStorage(ctx, fileStorageCfg)
	if err != nil {
		return nil, err
	}

	err = fs.migrateFileStorage(ctx)
	if err != nil {
		return nil, err
	}

	return &fs, nil
}

func (db *FileStorage) connectFileStorage(
	ctx context.Context,
	dbCfg Config, //nolint:gocritic // Its fine to pass config by value
) error {
	//nolint:mnd // No problem with these magic numbers
	fsConnectRetrier := retry.NewRetrier(
		retry.MaxRetries(5),
		retry.Delay(1*time.Second),
	)

	attempt := 0
	err := fsConnectRetrier.Wrap(
		func() error {
			attempt++
			db.logger.DebugContext(
				ctx,
				"connecting to FileStorage",
				logAttrs.Host(dbCfg.Host),
				logAttrs.Port(dbCfg.Port),
				logAttrs.Attempt(attempt),
			)

			var err error

			db.GormFileStorage, err = gorm.Open(
				postgres.Open(dbCfg.AsDSN()),
				&gorm.Config{}, //nolint:exhaustruct // TODO: inspect this cfg
			)
			if err != nil {
				db.logger.InfoContext(
					ctx,
					"failed to connect to FileStorage",
					logAttrs.Error(err),
					logAttrs.Attempt(attempt),
				)

				return fmt.Errorf("failed to connect to FileStorage: %w", err)
			}

			return nil
		},
	)()

	return err
}

func (db *FileStorage) migrateFileStorage(ctx context.Context) error {
	db.logger.DebugContext(ctx, "migrating FileStorage")

	err := db.GormFileStorage.AutoMigrate(
		gormModels.FileModel,
	)
	if err != nil {
		return fmt.Errorf("failed to migrate FileStorage: %w", err)
	}

	return nil
}

func (*FileStorage) Shutdown(_ context.Context) error {
	return nil
}
