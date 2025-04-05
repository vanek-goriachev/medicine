package chi

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"medicine/internal/appcore/collections"
	"medicine/internal/tooling/go-chi/middleware"
	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

type App struct {
	router   chi.Router
	logger   *slog.Logger
	mappers  *mappers
	services *chiServices
	server   *http.Server
	config   Config
}

func (r *App) Initialize(
	ctx context.Context,
	config Config,
	commonMappers *collections.CommonMappers,
	userActions *collections.UserActions,
	logger *slog.Logger,
) error {
	r.logger = logger
	r.mappers = newChiMappers(commonMappers)
	r.services = newChiServices(r.mappers, userActions)

	r.config = config

	err := r.initializeRouter(ctx)
	if err != nil {
		return err
	}

	r.initializeServer()

	return nil
}

func (r *App) initializeRouter(ctx context.Context) error {
	r.router = chi.NewRouter()

	r.registerMiddlewares()

	err := r.registerHandlers(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *App) registerMiddlewares() {
	r.router.Use(middleware.Authentification())
}

func (r *App) registerHandlers(ctx context.Context) error {
	openApiDefinition := generateApiSpec(r.services)

	err := openApiDefinition.SetupRoutes(r.router, openApiDefinition)
	if err != nil {
		r.logger.ErrorContext(ctx, "error setting up routes", logAttrs.Error(err))
		return fmt.Errorf("error setting up routes: %w", err)
	}

	return nil
}

func (r *App) initializeServer() {
	r.server = &http.Server{ //nolint:exhaustruct // TODO: inspect this cfg
		Addr:        fmt.Sprintf(":%d", r.config.Port),
		ReadTimeout: r.config.ReadTimeout,
		IdleTimeout: r.config.IdleTimeout,
		Handler:     r.router,
	}
}

func (r *App) Run(ctx context.Context) error {
	r.logger.InfoContext(
		ctx,
		"Starting server on port",
		logAttrs.Port(r.config.Port),
	)

	err := r.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error running server: %w", err)
	}

	return nil
}

func (r *App) Shutdown(ctx context.Context) error {
	r.logger.DebugContext(ctx, "Shutting down server")

	err := r.server.Shutdown(ctx)

	if err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	return nil
}
