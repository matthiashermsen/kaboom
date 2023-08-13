package logging

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/configuration"
)

func TestGetConfiguration(suite *testing.T) {
	suite.Run("Default log level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration.InitializeConfigurationEnvironment()

		config, err := GetConfiguration()

		assert.NoError(testing, err, "Expected no error when getting default configuration")

		assert.Equal(testing, zerolog.InfoLevel, config.Level, fmt.Sprintf("Expected log level to be %s", zerolog.InfoLevel))
	})

	suite.Run("Custom log level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration.InitializeConfigurationEnvironment()

		expectedLogLevel := zerolog.TraceLevel

		err := os.Setenv(configuration.EnvPrefix+"_LOGGING_LEVEL", strconv.Itoa(int(expectedLogLevel)))

		assert.NoError(testing, err, "Expected no error when setting environment variable")

		loggingConfiguration, err := GetConfiguration()

		assert.NoError(testing, err, "Expected no error when getting configuration")

		assert.Equal(testing, expectedLogLevel, loggingConfiguration.Level, fmt.Sprintf("Expected log level to be %d but got %d", expectedLogLevel, loggingConfiguration.Level))
	})

	suite.Run("Invalid log level", func(testing *testing.T) {
		defer os.Clearenv()

		configuration.InitializeConfigurationEnvironment()

		expectedLogLevel := 42

		err := os.Setenv(configuration.EnvPrefix+"_LOGGING_LEVEL", strconv.Itoa(expectedLogLevel))

		assert.NoError(testing, err, "Expected no error when setting environment variable")

		loggingConfiguration, err := GetConfiguration()

		assert.NoError(testing, err, "Expected no error when getting configuration")

		assert.Equal(testing, zerolog.Level(expectedLogLevel), loggingConfiguration.Level, fmt.Sprintf("Expected log level to be %d but got %d", expectedLogLevel, loggingConfiguration.Level))
	})
}
