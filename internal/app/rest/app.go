package rest

import (
	"context"
	"fmt"
	"syscall"

	"medicine/internal/app/rest/chi"
	"medicine/internal/appcore"
	logAttrs "medicine/pkg/telemetry/logging/logging-attributes"
)

type App struct {
	systemDependencies *appcore.SystemDependencies
	appCore            *appcore.Core
	chiApp             *chi.App
}

func NewApp(systemDependencies *appcore.SystemDependencies) *App {
	return &App{
		systemDependencies: systemDependencies, // systemDependencies are duplicated in the appCore

		//nolint:exhaustruct // Will be initialized in .initialize() method
		appCore: &appcore.Core{},
		chiApp:  &chi.App{},
	}
}

func (a *App) InitializeAndRun() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := a.initialize(ctx)
	if err != nil {
		panic(fmt.Errorf("initialization failed: %w", err))
	}

	go a.cancelContextOnReceiveInterruptingSignals(ctx, cancel)

	err = a.runWithGracefulShutdown(ctx)
	if err != nil {
		panic(fmt.Errorf("running failed: %w", err))
	}
}

func (a *App) cancelContextOnReceiveInterruptingSignals(ctx context.Context, cancelFunc context.CancelFunc) {
	terminal := a.systemDependencies.Terminal
	signal := <-terminal.SignalCh(1, syscall.SIGINT, syscall.SIGTERM)
	a.logger().InfoContext(ctx, "Received interrupting signal. Shutting down...", logAttrs.Signal(signal.String()))
	cancelFunc()
}
