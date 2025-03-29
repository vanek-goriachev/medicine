package rest

//nolint:wrapcheck // Error will be wrapped in the caller
func (a *App) printStringToStdout(s string) (int, error) {
	return a.systemDependencies.Terminal.Stdout().Write([]byte(s))
}
