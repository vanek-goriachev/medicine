package appcore

import (
	"context"
	"fmt"
	"log/slog"
	noop_authorizer "medicine/internal/layers/business-logic/authorization/noop-authorizer"

	"medicine/internal/appcore/collections"
	"medicine/internal/appcore/dependencies"
)

type Core struct {
	systemDependencies      *SystemDependencies
	ApplicationDependencies *dependencies.ApplicationDependencies

	others *collections.Others

	CommonMappers *collections.CommonMappers

	dbMappers  *collections.DBMappers
	dbGateways *collections.DBGateways

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

	c.dbMappers = collections.NewDBMappers()
	c.dbGateways = collections.NewDBGateways(c.ApplicationDependencies.DB, c.dbMappers)

	c.validators = collections.NewValidators()
	c.factories = collections.NewFactories(c.validators)

	authorizer := noop_authorizer.NewNoopAuthorizer(
		c.ApplicationDependencies.DB.GormDB,
	)
	c.simpleActions = collections.NewSimpleActions(authorizer, c.others, c.dbGateways, c.factories)

	c.UserActions = collections.NewUserActions(
		authorizer,
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
