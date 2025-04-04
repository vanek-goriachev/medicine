package logging_attributes

import "log/slog"

func Attempt(attempt int) slog.Attr {
	return slog.Int(
		"attempt",
		attempt,
	)
}
