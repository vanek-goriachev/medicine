package chi

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"medicine/internal/appcore/collections"
	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

type App struct {
	router   chi.Router
	logger   *slog.Logger
	mappers  *mappers
	services *services
	server   *http.Server
	config   Config
}

func (r *App) Initialize(
	config Config,
	commonMappers *collections.CommonMappers,
	userActions *collections.UserActions,
	logger *slog.Logger,
) {
	r.logger = logger
	r.mappers = newChiMappers(commonMappers)
	r.services = newChiServices(r.mappers, userActions)

	r.config = config
	r.initializeRouter()
	r.initializeServer()
}

func (r *App) initializeRouter() {
	r.router = chi.NewRouter()

	r.router.Post("/api/v1/tags-space/", r.services.tagsSpace.Create)
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
