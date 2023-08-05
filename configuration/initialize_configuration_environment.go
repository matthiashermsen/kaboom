package configuration

import "github.com/spf13/viper"

const EnvPrefix = "KABOOM"

func InitializeConfigurationEnvironment() {
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()
}
