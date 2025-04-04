package dependencies

import (
	"context"
	"errors"
	"fmt"

	"medicine/internal/appcore/dependencies/db"
	"medicine/pkg/telemetry"
	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

type ApplicationDependencies struct {
	Telemetry telemetry.Infrastructure
	// IAM
	DB *db.DB
}

func (a *ApplicationDependencies) Initialize(
	ctx context.Context,
	depsCfg DepsConfig, //nolint:gocritic // Its fine to pass config by value
) error {
	var err error

	a.Telemetry, err = telemetry.InitInfrastructure(ctx, depsCfg.Telemetry)
	if err != nil {
		return fmt.Errorf("failed to initialize telemetry: %w", err)
	}

	a.Telemetry.Logging.Logger.DebugContext(ctx, "Initialized telemetry")

	// Initialize IAM

	a.Telemetry.Logging.Logger.DebugContext(ctx, "Initializing DB")

	a.DB, err = db.NewDB(ctx, depsCfg.DB, a.Telemetry.Logging.Logger)
	if err != nil {
		// Log this error here, because on upper levels we can't be sure that logging was initialized correctly
		a.Telemetry.Logging.Logger.ErrorContext(ctx, "Failed to initialize DB", logAttrs.Error(err))

		return fmt.Errorf("failed to initialize DB: %w", err)
	}

	return nil
}

func (a *ApplicationDependencies) Shutdown(ctx context.Context) error {
	a.Telemetry.Logging.Logger.DebugContext(ctx, "Shutting down DB")

	dbShutdownErr := a.DB.Shutdown(ctx)

	a.Telemetry.Logging.Logger.DebugContext(ctx, "Shutting down telemetry")

	telemetryShutdownErr := a.Telemetry.Shutdown(ctx)

	var shutdownErrors = errors.Join(dbShutdownErr, telemetryShutdownErr)

	return shutdownErrors
}
