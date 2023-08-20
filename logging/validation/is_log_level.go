package validation

import (
	"log/slog"
	"slices"
)

func IsLogLevel(value interface{}) bool {
	logLevel, isLogLevel := value.(slog.Level)

	if !isLogLevel {
		return false
	}

	logLevels := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}

	return slices.Contains(logLevels, logLevel)
}
