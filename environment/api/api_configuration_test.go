package api_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/environment"
	"github.com/matthiashermsen/kaboom/environment/api"
)

func TestNewAPIConfiguration(suite *testing.T) {
	suite.Run("Custom port", func(testing *testing.T) {
		defer os.Clearenv()

		expectedPort := 8080

		err := os.Setenv(environment.EnvPrefix+"_"+api.EnvPort, strconv.Itoa(expectedPort))
		assert.NoError(testing, err, "Expected no error when setting server port")

		configuration, err := api.NewAPIConfiguration()
		assert.NoError(testing, err, "Expected no error when creating configuration")

		assert.Equal(testing, expectedPort, configuration.Port, fmt.Sprintf("Expected API port to be %v", expectedPort))
	})

	suite.Run("Default port", func(testing *testing.T) {
		defer os.Clearenv()

		configuration, err := api.NewAPIConfiguration()
		assert.NoError(testing, err, "Expected no error when creating configuration")

		assert.Equal(testing, api.DefaultPort, configuration.Port, fmt.Sprintf("Expected API port to be %v", api.DefaultPort))
	})

	suite.Run("Invalid port", func(testing *testing.T) {
		defer os.Clearenv()

		err := os.Setenv(environment.EnvPrefix+"_"+api.EnvPort, "0")
		assert.NoError(testing, err, "Expected no error when setting server port")

		_, err = api.NewAPIConfiguration()

		assert.Error(testing, err, "Expected error for invalid port")
	})
}
