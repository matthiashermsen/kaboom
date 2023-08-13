package logging

import (
	"fmt"
	"log/slog"
)

type LogLevelInvalidError struct{}

func (logLevelInvalidError LogLevelInvalidError) Error() string {
	return fmt.Sprintf("Log level is invalid. Must be one of [%d, %d, %d, %d].", slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError)
}
