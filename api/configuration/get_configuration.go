package configuration

import "github.com/spf13/viper"

func GetConfiguration() (Configuration, error) {
	viper.SetDefault("SERVER_PORT", 3000)

	var configuration Configuration

	err := viper.Unmarshal(&configuration)

	return configuration, err
}
