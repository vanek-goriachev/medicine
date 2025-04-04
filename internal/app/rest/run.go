package rest

import (
	"context"
	"errors"
	"fmt"

	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

func (a *App) runWithGracefulShutdown(ctx context.Context) error {
	a.logger().DebugContext(ctx, "Running application")

	go a.gracefulShutdowns(ctx)

	err := a.chiApp.Run(ctx)

	return fmt.Errorf("error while running the server: %w", err)
}

func (a *App) gracefulShutdowns(ctx context.Context) {
	<-ctx.Done()

	a.logger().InfoContext(ctx, "Shutting down the chi application (rest)")

	chiShutdownErr := a.chiApp.Shutdown(ctx)
	if chiShutdownErr != nil {
		a.logger().ErrorContext(
			ctx,
			"error while shutting down the chi application",
			logAttrs.Error(chiShutdownErr),
		)
	}

	a.logger().InfoContext(ctx, "Shutting down the app core")

	appCoreShutdownErr := a.appCore.Shutdown(ctx)

	var shutdownErrors = errors.Join(chiShutdownErr, appCoreShutdownErr)

	if shutdownErrors != nil {
		//nolint:errcheck // Cant do anything about this error
		_, _ = a.toStderr("shutdown errors: " + shutdownErrors.Error())
	}
}
