package logging_attributes

import "log/slog"

func Host(host string) slog.Attr {
	return slog.String("host", host)
}

func Port(port int) slog.Attr {
	return slog.Int("port", port)
}

func Signal(signal string) slog.Attr {
	return slog.String("signal", signal)
}
