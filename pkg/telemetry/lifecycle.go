package telemetry

import (
	"context"
	"errors"
	"fmt"

	"medicine/pkg/telemetry/logging"
)

func InitInfrastructure(ctx context.Context, cfg Config) (*Infrastructure, error) {
	loggingInfrastructure, err := logging.InitInfrastructure(cfg.Logging)
	if err != nil {
		return nil, fmt.Errorf("error on logging initialization: %w", err)
	}

	loggingInfrastructure.Logger.DebugContext(ctx, "Started logging")

	return &Infrastructure{
		Logging: loggingInfrastructure,
	}, nil
}

func (i *Infrastructure) Shutdown(ctx context.Context) error {
	i.Logging.Logger.DebugContext(ctx, "Shutting down logging infrastructure")

	loggingErr := i.Logging.Shutdown(ctx)

	shutdownErrors := errors.Join(loggingErr)
	if shutdownErrors != nil {
		return fmt.Errorf("error on shutting down logging infrastructure: %w", shutdownErrors)
	}

	return nil
}
