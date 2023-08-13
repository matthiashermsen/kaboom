package logging

import "log/slog"

func ValidateConfiguration(configuration Configuration) error {
	logLevels := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}

	for _, logLevel := range logLevels {
		if configuration.Level == logLevel {
			return nil
		}
	}

	return LogLevelInvalidError{}
}
