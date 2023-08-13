package logging

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateConfiguration(suite *testing.T) {
	suite.Run("No error", func(suite *testing.T) {
		logLevels := []slog.Level{
			slog.LevelDebug,
			slog.LevelInfo,
			slog.LevelWarn,
			slog.LevelError,
		}

		for _, logLevel := range logLevels {
			suite.Run(fmt.Sprintf("Log level=%s", logLevel), func(testing *testing.T) {
				configuration := Configuration{Level: logLevel}

				err := ValidateConfiguration(configuration)

				assert.NoError(testing, err, "Expected no validation error")
			})
		}
	})

	suite.Run("LogLevelInvalidError", func(suite *testing.T) {
		logLevels := []int{
			-10,
			-5,
			5,
			10,
			1,
			-1,
		}

		for _, logLevel := range logLevels {
			suite.Run(fmt.Sprintf("Log level=%d", logLevel), func(testing *testing.T) {
				configuration := Configuration{Level: slog.Level(logLevel)}

				err := ValidateConfiguration(configuration)

				assert.Error(testing, err, "Expected validation error")
				assert.IsType(testing, LogLevelInvalidError{}, err, "Expected LogLevelInvalidError")
			})
		}
	})
}
