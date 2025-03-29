package rest

import (
	"context"
)

func (a *App) run(ctx context.Context) error { //nolint:revive // going to fix later
	a.printStringToStdout("Running...\n") //nolint:revive,errcheck // going to remove later

	// Запустить сервер
	// Настроить Graceful Shutdown (в том числе телеметрии)

	return nil
}
