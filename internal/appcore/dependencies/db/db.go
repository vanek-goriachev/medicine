package db

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormModels "medicine/internal/layers/storage/db/gorm/models"
	"medicine/pkg/retry"
	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

type DB struct {
	GormDB *gorm.DB
	logger *slog.Logger
}

func NewDB(
	ctx context.Context,
	dbCfg Config, //nolint:gocritic // Its fine to pass config by value
	logger *slog.Logger,
) (*DB, error) {
	var db DB
	db.logger = logger

	err := db.connectDB(ctx, dbCfg)
	if err != nil {
		return nil, err
	}

	err = db.migrateDB(ctx)
	if err != nil {
		return nil, err
	}

	return &db, nil
}

func (db *DB) connectDB(
	ctx context.Context,
	dbCfg Config, //nolint:gocritic // Its fine to pass config by value
) error {
	//nolint:mnd // No problem with these magic numbers
	dbConnectRetrier := retry.NewRetrier(
		retry.MaxRetries(5),
		retry.Delay(1*time.Second),
	)

	attempt := 0
	err := dbConnectRetrier.Wrap(
		func() error {
			attempt++
			db.logger.DebugContext(
				ctx,
				"connecting to DB",
				logAttrs.Host(dbCfg.Host),
				logAttrs.Port(dbCfg.Port),
				logAttrs.Attempt(attempt),
			)

			var err error

			db.GormDB, err = gorm.Open(
				postgres.Open(dbCfg.AsDSN()),
				&gorm.Config{}, //nolint:exhaustruct // TODO: inspect this cfg
			)
			if err != nil {
				db.logger.InfoContext(
					ctx,
					"failed to connect to DB",
					logAttrs.Error(err),
					logAttrs.Attempt(attempt),
				)

				return fmt.Errorf("failed to connect to DB: %w", err)
			}

			return nil
		},
	)()

	return err
}

func (db *DB) migrateDB(ctx context.Context) error {
	db.logger.DebugContext(ctx, "migrating DB")

	err := db.GormDB.AutoMigrate(
		gormModels.TagsSpaceModel,
		gormModels.TagModel,
	)
	if err != nil {
		return fmt.Errorf("failed to migrate DB: %w", err)
	}

	return nil
}

func (*DB) Shutdown(_ context.Context) error {
	return nil
}
