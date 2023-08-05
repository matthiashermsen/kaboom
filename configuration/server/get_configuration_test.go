package server

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/configuration"
)

func TestGetConfiguration_DefaultPort(testing *testing.T) {
	defer os.Clearenv()

	configuration.InitializeConfigurationEnvironment()

	config, err := GetConfiguration()

	assert.NoError(testing, err, "Expected no error when getting default configuration")

	var expectedPort uint16 = 3000

	assert.Equal(testing, expectedPort, config.Port, fmt.Sprintf("Expected port to be %v", expectedPort))
}

func TestGetConfiguration_CustomPort(testing *testing.T) {
	defer os.Clearenv()

	configuration.InitializeConfigurationEnvironment()

	var expectedPort uint16 = 8080

	err := os.Setenv(configuration.EnvPrefix+"_SERVER_PORT", strconv.Itoa(int(expectedPort)))

	assert.NoError(testing, err, "Expected no error when setting environment variable")

	config, err := GetConfiguration()

	assert.NoError(testing, err, "Expected no error when getting configuration")

	assert.Equal(testing, expectedPort, config.Port, fmt.Sprintf("Expected port to be %d but got %d", expectedPort, config.Port))
}

func TestGetConfiguration_InvalidPort(testing *testing.T) {
	defer os.Clearenv()

	configuration.InitializeConfigurationEnvironment()

	expectedPort := "made-up"

	err := os.Setenv(configuration.EnvPrefix+"_SERVER_PORT", expectedPort)

	assert.NoError(testing, err, "Expected no error when setting environment variable")

	_, err = GetConfiguration()

	assert.Error(testing, err, "Expected error when getting configuration")
}
