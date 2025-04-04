package logging

import (
	"log/slog"
)

type Infrastructure struct {
	Logger *slog.Logger
}

func (i *Infrastructure) GetLogger() *slog.Logger {
	return i.Logger
}
