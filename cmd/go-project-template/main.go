package main

import (
	"template/internal/app/rest"
	"template/pkg/os"
)

func main() {
	terminal := os.NewTerminal()

	_, err := terminal.Stdout().Write([]byte("Creating new application\n"))
	if err != nil {
		return
	}

	app := rest.NewApp(
		terminal,
	)

	_, err = terminal.Stdout().Write([]byte("Initializing and running the application\n"))
	if err != nil {
		return
	}

	err = app.InitializeAndRun()
	if err != nil {
		//nolint:errcheck // We cant do anything if write failed
		_, _ = terminal.Stderr().Write([]byte(err.Error()))
	}
}
