package os

import (
	"io"
	"os/signal"

	sysos "os"
)

type Terminal struct{}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (*Terminal) ExitWithCode(code int) {
	sysos.Exit(code)
}

func (*Terminal) Stdout() io.Writer {
	return sysos.Stdout
}

func (*Terminal) Stderr() io.Writer {
	return sysos.Stderr
}

func (*Terminal) SignalCh(size int, signals ...sysos.Signal) <-chan sysos.Signal {
	signalCh := make(chan sysos.Signal, size)
	signal.Notify(signalCh, signals...)

	return signalCh
}
