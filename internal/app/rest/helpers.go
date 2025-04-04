package rest

import "log/slog"

//nolint:wrapcheck // Error will be wrapped in the caller
func (a *App) toStdout(s string) (int, error) {
	return a.systemDependencies.Terminal.Stdout().Write([]byte(s))
}

//nolint:wrapcheck // Error will be wrapped in the caller if needed
func (a *App) toStderr(s string) (int, error) {
	return a.systemDependencies.Terminal.Stderr().Write([]byte(s))
}

func (a *App) logger() *slog.Logger {
	return a.appCore.ApplicationDependencies.Telemetry.Logging.Logger
}
