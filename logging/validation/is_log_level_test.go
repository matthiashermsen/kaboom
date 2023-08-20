package validation

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLogLevel(suite *testing.T) {
	suite.Run("Valid log levels", func(suite *testing.T) {
		validLogLevels := []slog.Level{
			slog.LevelDebug,
			slog.LevelInfo,
			slog.LevelWarn,
			slog.LevelError,
		}

		for _, logLevel := range validLogLevels {
			suite.Run(fmt.Sprintf("Log level=%s", logLevel), func(testing *testing.T) {
				assert.True(testing, IsLogLevel(logLevel), "Expected IsLogLevel to return true")
			})
		}
	})

	suite.Run("Invalid log levels", func(suite *testing.T) {
		type CustomType struct {
			Value int
		}

		invalidLogLevels := []interface{}{
			"slog.LevelDebug",
			123,
			3.14,
			true,
			nil,
			CustomType{Value: 42},
		}

		for _, logLevel := range invalidLogLevels {
			suite.Run(fmt.Sprintf("Log level=%v", logLevel), func(testing *testing.T) {
				assert.False(testing, IsLogLevel(logLevel), "Expected IsLogLevel to return false for value %v", logLevel)
			})
		}
	})
}
