package appcore

import (
	"context"
	"fmt"
	"log/slog"

	"medicine/internal/appcore/collections"
	"medicine/internal/appcore/dependencies"
)

type Core struct {
	systemDependencies      *SystemDependencies
	ApplicationDependencies *dependencies.ApplicationDependencies

	others *collections.Others

	CommonMappers *collections.CommonMappers
	gormMappers   *collections.GORMMappers

	gateways *collections.Gateways

	factories *collections.Factories

	// AtomicActions are implemented by gateways now
	simpleActions *collections.SimpleActions
	UserActions   *collections.UserActions
}

func (c *Core) Initialize(
	systemDependencies *SystemDependencies,
	applicationDependencies *dependencies.ApplicationDependencies,
) {
	c.systemDependencies = systemDependencies
	c.ApplicationDependencies = applicationDependencies

	c.others = collections.NewOthers()

	c.CommonMappers = collections.NewCommonMappers()
	c.gormMappers = collections.NewGORMMappers()

	c.gateways = collections.NewGateways(c.ApplicationDependencies.DB.GormDB, c.gormMappers)

	c.factories = collections.NewFactories()

	c.simpleActions = collections.NewSimpleActions(c.others, c.gateways, c.factories)
	c.UserActions = collections.NewUserActions(c.simpleActions)
}

func (c *Core) Shutdown(ctx context.Context) error {
	c.logger().DebugContext(ctx, "Shutting down the appcore")

	err := c.ApplicationDependencies.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to shutdown the appcore: %w", err)
	}

	return nil
}

func (c *Core) logger() *slog.Logger {
	return c.ApplicationDependencies.Telemetry.Logging.Logger
}
