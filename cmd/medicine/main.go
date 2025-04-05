package main

import (
	"fmt"

	"medicine/internal/app/rest"
	"medicine/internal/appcore"
	"medicine/pkg/os"
)

func main() {
	fmt.Println("Running application") //nolint:forbidigo // Print is fine here and no need to handle error

	terminal := os.NewTerminal()
	systemDependencies := &appcore.SystemDependencies{
		Terminal: terminal,
	}

	app := rest.NewApp(systemDependencies)

	app.InitializeAndRun()
}
