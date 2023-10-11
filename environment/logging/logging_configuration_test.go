package logging_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/environment"
	"github.com/matthiashermsen/kaboom/environment/logging"
)

func TestNewLoggingConfiguration(suite *testing.T) {
	suite.Run("Default level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration, err := logging.NewLoggingConfiguration()
		assert.NoError(testing, err, "Expected no error when creating configuration")

		assert.Equal(testing, logging.DefaultLevel, configuration.Level, fmt.Sprintf("Expected level to be %v", logging.DefaultLevel))
	})

	suite.Run("Invalid level", func(testing *testing.T) {
		defer os.Clearenv()

		err := os.Setenv(environment.EnvPrefix+"_"+logging.EnvLevel, "made-up")
		assert.NoError(testing, err, "Expected no error when setting log level")

		_, err = logging.NewLoggingConfiguration()

		assert.Error(testing, err, "Expected error for invalid level")
	})

	suite.Run("Custom level", func(suite *testing.T) {
		levelMap := map[string]slog.Level{
			"debug": slog.LevelDebug,
			"info":  slog.LevelInfo,
			"warn":  slog.LevelWarn,
			"error": slog.LevelError,
		}

		for levelAsString, expectedLogLevel := range levelMap {
			suite.Run("Log level="+levelAsString, func(testing *testing.T) {
				defer os.Clearenv()

				err := os.Setenv(environment.EnvPrefix+"_"+logging.EnvLevel, levelAsString)
				assert.NoError(testing, err, "Expected no error when setting log level")

				configuration, err := logging.NewLoggingConfiguration()
				assert.NoError(testing, err, "Expected no error when creating configuration")

				assert.Equal(testing, expectedLogLevel, configuration.Level, fmt.Sprintf("Expected log level to be %s", expectedLogLevel))
			})
		}
	})
}
