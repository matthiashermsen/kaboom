package logging

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func GetConfiguration() (Configuration, error) {
	viper.SetDefault("LOGGING_LEVEL", zerolog.InfoLevel)

	var configuration Configuration

	err := viper.Unmarshal(&configuration)

	return configuration, err
}
