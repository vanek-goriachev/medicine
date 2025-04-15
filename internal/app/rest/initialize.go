package rest

import (
	"context"
	"fmt"
	"time"

	"medicine/internal/app/rest/chi"
	"medicine/internal/appcore/dependencies"
	"medicine/internal/appcore/dependencies/db"
	file_storage "medicine/internal/appcore/dependencies/file-storage"
	"medicine/internal/tooling/iam"
	"medicine/pkg/telemetry"
	"medicine/pkg/telemetry/logging"
)

const writeFail = "failed to write to stdout: %w"

func (a *App) initialize(ctx context.Context) error {
	// No logger here since it is not initialized yet
	_, err := a.toStdout("Initializing application\n")
	if err != nil {
		return fmt.Errorf(writeFail, err)
	}

	// No logger here since it is not initialized yet
	_, err = a.toStdout("Loading filesystem data\n")
	if err != nil {
		return fmt.Errorf(writeFail, err)
	}

	restCfg, coreDepsCfg, err := a.loadFileSystemData()
	if err != nil {
		return fmt.Errorf("failed to load filesystem data: %w", err)
	}

	var applicationDependencies dependencies.ApplicationDependencies

	err = applicationDependencies.Initialize(ctx, coreDepsCfg) // Ping and migrations happens here
	if err != nil {
		return fmt.Errorf("failed to initialize dependencies: %w", err)
	}

	// Don't use a.logger() here because it refers to appcore which is not initialized yet
	applicationDependencies.Telemetry.Logging.Logger.DebugContext(ctx, "Initializing appcore (business logic)")
	a.appCore.Initialize(a.systemDependencies, &applicationDependencies)

	a.logger().DebugContext(ctx, "Initializing chi-application (rest)")

	err = a.chiApp.Initialize(
		ctx,
		restCfg,
		a.appCore.CommonMappers,
		a.appCore.UserActions,
		a.logger(),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize chi-application (rest): %w", err)
	}

	return nil
}

//nolint:mnd,unparam // Hardcode will be removed later
func (*App) loadFileSystemData() (chi.Config, dependencies.DepsConfig, error) {
	restCfg := chi.Config{
		Port:        8080,
		ReadTimeout: time.Millisecond * 100,
		IdleTimeout: time.Second * 30,
	}

	coreDepsCfg := dependencies.DepsConfig{
		DB: db.Config{
			Host:     "db",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DBName:   "postgres",
			SslMode:  "disable",
			Timezone: "UTC",
		},
		FileStorage: file_storage.Config{
			Host:     "file-storage",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DBName:   "postgres",
			SslMode:  "disable",
			Timezone: "UTC",
		},
		Telemetry: telemetry.Config{
			Logging: logging.Config{},
		},
		IAM: iam.Config{},
	}

	return restCfg, coreDepsCfg, nil
}
