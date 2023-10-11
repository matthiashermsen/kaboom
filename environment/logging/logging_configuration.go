package logging

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const EnvLevel = "LOGGING_LEVEL"
const DefaultLevel = slog.LevelInfo

type LoggingConfiguration struct {
	Level slog.Level
}

func NewLoggingConfiguration() (*LoggingConfiguration, error) {
	viper.AutomaticEnv()

	level, err := getLevel()

	if err != nil {
		return nil, err
	}

	loggingConfiguration := LoggingConfiguration{Level: level}
	loggingConfigurationValidationError := loggingConfiguration.validate()

	return &loggingConfiguration, loggingConfigurationValidationError
}

func getLevel() (slog.Level, error) {
	levelAsString := viper.GetString(EnvLevel)
	lowerCaseLevelAsString := strings.ToLower(levelAsString)

	switch lowerCaseLevelAsString {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	case "":
		return DefaultLevel, nil
	default:
		return DefaultLevel, fmt.Errorf("invalid log level '%s'", levelAsString)
	}
}

func (loggingConfiguration *LoggingConfiguration) validate() error {
	validator := validator.New()
	err := validator.Struct(loggingConfiguration)

	return err
}
