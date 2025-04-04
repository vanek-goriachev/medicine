package main

import (
	"medicine/internal/app/rest"
	"medicine/internal/appcore"
	"medicine/pkg/os"
)

func main() {
	terminal := os.NewTerminal()
	systemDependencies := &appcore.SystemDependencies{
		Terminal: terminal,
	}

	app := rest.NewApp(systemDependencies)

	app.InitializeAndRun()
}
