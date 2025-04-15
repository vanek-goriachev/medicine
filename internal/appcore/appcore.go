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

	dbMappers  *collections.DBMappers
	dbGateways *collections.DBGateways

	fileStorageMappers  *collections.FileStorageMappers
	fileStorageGateways *collections.FileStorageGateways

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

	c.fileStorageMappers = collections.NewFileStorageMappers()
	c.fileStorageGateways = collections.NewFileStorageGateways(
		c.ApplicationDependencies.FileStorage,
		c.fileStorageMappers,
	)

	c.validators = collections.NewValidators()
	c.factories = collections.NewFactories(c.validators)

	c.simpleActions = collections.NewSimpleActions(c.others, c.dbGateways, c.factories)

	authorizer := authorization.NewAuthorizer(c.ApplicationDependencies.IAM)
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
