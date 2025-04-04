package logging

import (
	"context"
	"log/slog"
)

func InitInfrastructure(
	_ Config,
) (Infrastructure, error) { //nolint:unparam // Will be fixed later
	logger := slog.Default()

	return Infrastructure{
		Logger: logger,
	}, nil
}

func (*Infrastructure) Shutdown(_ context.Context) error {
	return nil
}
