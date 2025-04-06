package appcore

import (
	"context"
	"fmt"
	"log/slog"

	"medicine/internal/appcore/collections"
	"medicine/internal/appcore/dependencies"
	"medicine/internal/layers/business-logic/authorization"
)

type Core struct {
	systemDependencies      *SystemDependencies
	ApplicationDependencies *dependencies.ApplicationDependencies

	others *collections.Others

	CommonMappers *collections.CommonMappers
	gormMappers   *collections.GORMMappers

	gateways *collections.Gateways

	validators *collections.Validators
	factories  *collections.Factories

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

	c.validators = collections.NewValidators()
	c.factories = collections.NewFactories(c.validators)

	c.simpleActions = collections.NewSimpleActions(c.others, c.gateways, c.factories)
	c.UserActions = collections.NewUserActions(
		authorization.NewAuthorizer(c.ApplicationDependencies.IAM),
		c.simpleActions,
	)
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
