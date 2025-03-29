package rest

import (
	"io"
	"os"
)

type SystemDependencies struct {
	// Terminal is an interface that provides access to terminal functions.
	Terminal Terminal

	// Clock
	// Filesystem
}

type Terminal interface {
	ExitWithCode(code int)
	Stdout() io.Writer
	Stderr() io.Writer
	// SignalCh returns a channel that will receive the specified signals.
	SignalCh(size int, signals ...os.Signal) <-chan os.Signal
}
