package logging

import "log/slog"

type Configuration struct {
	Level slog.Level `mapstructure:"LOGGING_LEVEL"`
}
