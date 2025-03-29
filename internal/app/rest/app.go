package rest

import (
	"context"
	"fmt"
	"syscall"

	"template/internal/appcore"
)

type App struct {
	systemDependencies *SystemDependencies
	appCore            *appcore.Core
}

func NewApp(terminal Terminal) *App {
	return &App{
		systemDependencies: &SystemDependencies{
			Terminal: terminal,
		},
		appCore: nil, // Will be initialized in .initialize() method
	}
}

func (a *App) InitializeAndRun() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := a.initialize(ctx)
	if err != nil {
		return fmt.Errorf("initialization failed: %w", err)
	}

	go a.cancelContextOnReceiveInterruptingSignals(cancel)

	err = a.run(ctx)
	if err != nil {
		return fmt.Errorf("running failed: %w", err)
	}

	return nil
}

func (a *App) cancelContextOnReceiveInterruptingSignals(cancelFunc context.CancelFunc) {
	terminal := a.systemDependencies.Terminal
	<-terminal.SignalCh(1, syscall.SIGINT, syscall.SIGTERM)
	cancelFunc()
}
