package rest

import (
	"context"
	"fmt"
	"time"

	"medicine/internal/app/rest/chi"
	"medicine/internal/appcore/dependencies"
	"medicine/internal/appcore/dependencies/db"
	"medicine/pkg/telemetry"
	"medicine/pkg/telemetry/logging"
)

const writeFail = "failed to write to stdout: %w"

func (a *App) initialize(ctx context.Context) error {
	_, err := a.toStdout("Initializing application\n") // No logger here since it is not initialized yet
	if err != nil {
		return fmt.Errorf(writeFail, err)
	}

	_, err = a.toStdout("Loading filesystem data\n") // No logger here since it is not initialized yet
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

	applicationDependencies.Telemetry.Logging.Logger.DebugContext(ctx, "Initializing appcore (business logic)")
	a.appCore.Initialize(a.systemDependencies, &applicationDependencies)

	applicationDependencies.Telemetry.Logging.Logger.DebugContext(ctx, "Initializing chi-application (rest)")
	a.chiApp.Initialize(
		restCfg,
		a.appCore.CommonMappers,
		a.appCore.UserActions,
		applicationDependencies.Telemetry.Logging.Logger,
	)

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
		Telemetry: telemetry.Config{
			Logging: logging.Config{},
		},
		IAM: dependencies.IAMConfig{},
	}

	return restCfg, coreDepsCfg, nil
}
