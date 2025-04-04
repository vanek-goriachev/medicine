package rest

import (
	"context"
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
		_, err = a.toStderr("initialization failed: " + err.Error())
		panic(err)
	}

	go a.cancelContextOnReceiveInterruptingSignals(ctx, cancel)

	err = a.runWithGracefulShutdown(ctx)
	if err != nil {
		_, err = a.toStderr("running failed: " + err.Error())
		panic(err)
	}
}

func (a *App) cancelContextOnReceiveInterruptingSignals(ctx context.Context, cancelFunc context.CancelFunc) {
	terminal := a.systemDependencies.Terminal
	signal := <-terminal.SignalCh(1, syscall.SIGINT, syscall.SIGTERM)
	a.appCore.ApplicationDependencies.Telemetry.Logging.Logger.InfoContext(
		ctx,
		"Received interrupting signal. Shutting down...",
		logAttrs.Signal(signal.String()),
	)
	cancelFunc()
}
