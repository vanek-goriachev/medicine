package logging_attributes

import "log/slog"

func Error(err error) slog.Attr {
	return slog.String("error", err.Error())
}
