package environment

import (
	"github.com/spf13/viper"

	"github.com/matthiashermsen/kaboom/environment/api"
	"github.com/matthiashermsen/kaboom/environment/logging"
)

const EnvPrefix = "KABOOM"

func init() {
	viper.SetEnvPrefix(EnvPrefix)
}

type Configuration struct {
	Logging logging.LoggingConfiguration
	API     api.APIConfiguration
}

func NewConfiguration() (*Configuration, error) {
	viper.AutomaticEnv()

	loggingConfiguration, err := logging.NewLoggingConfiguration()

	if err != nil {
		return nil, err
	}

	apiConfiguration, err := api.NewAPIConfiguration()

	if err != nil {
		return nil, err
	}

	configuration := Configuration{
		Logging: *loggingConfiguration,
		API:     apiConfiguration,
	}

	return &configuration, nil
}
