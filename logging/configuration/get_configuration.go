package configuration

import (
	"log/slog"

	"github.com/spf13/viper"
)

func GetConfiguration() (Configuration, error) {
	viper.SetDefault("LOGGING_LEVEL", slog.LevelInfo)

	var configuration Configuration

	err := viper.Unmarshal(&configuration)

	return configuration, err
}
