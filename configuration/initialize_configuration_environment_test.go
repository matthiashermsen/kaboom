package configuration

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitializeConfigurationEnvironment(testing *testing.T) {
	defer os.Clearenv()

	environmentVariableKey := "FOO"
	expectedValue := "foo"

	err := os.Setenv(EnvPrefix+"_"+environmentVariableKey, expectedValue)

	assert.NoError(testing, err, "Expected no error when setting env variable")

	InitializeConfigurationEnvironment()

	assert.Equal(testing, expectedValue, viper.GetString(environmentVariableKey))
}
