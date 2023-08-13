package logging

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/configuration"
)

func TestGetConfiguration(suite *testing.T) {
	suite.Run("Default log level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration.InitializeConfigurationEnvironment()

		config, err := GetConfiguration()

		assert.NoError(testing, err, "Expected no error when getting default configuration")

		assert.Equal(testing, slog.LevelInfo, config.Level, fmt.Sprintf("Expected log level to be %s", slog.LevelInfo))
	})

	suite.Run("Custom log level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration.InitializeConfigurationEnvironment()

		expectedLogLevel := slog.LevelDebug

		err := os.Setenv(configuration.EnvPrefix+"_LOGGING_LEVEL", strconv.Itoa(int(expectedLogLevel)))

		assert.NoError(testing, err, "Expected no error when setting environment variable")

		loggingConfiguration, err := GetConfiguration()

		assert.NoError(testing, err, "Expected no error when getting configuration")

		assert.Equal(testing, expectedLogLevel, loggingConfiguration.Level, fmt.Sprintf("Expected log level to be %d but got %d", expectedLogLevel, loggingConfiguration.Level))
	})

	suite.Run("Invalid log level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration.InitializeConfigurationEnvironment()

		expectedLogLevel := 999

		err := os.Setenv(configuration.EnvPrefix+"_LOGGING_LEVEL", strconv.Itoa(expectedLogLevel))

		assert.NoError(testing, err, "Expected no error when setting environment variable")

		loggingConfiguration, err := GetConfiguration()

		assert.NoError(testing, err, "Expected no error when getting configuration")

		assert.Equal(testing, slog.Level(expectedLogLevel), loggingConfiguration.Level, fmt.Sprintf("Expected log level to be %d but got %d", expectedLogLevel, loggingConfiguration.Level))
	})
}
