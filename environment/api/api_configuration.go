package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const EnvPort = "API_PORT"
const DefaultPort = 3000

type APIConfiguration struct {
	Port int `validate:"gte=1,lte=65535"`
}

func NewAPIConfiguration() (APIConfiguration, error) {
	viper.AutomaticEnv()

	port := getPort()

	apiConfiguration := APIConfiguration{Port: port}
	apiConfigurationValidationError := apiConfiguration.validate()

	return apiConfiguration, apiConfigurationValidationError
}

func getPort() int {
	viper.SetDefault(EnvPort, DefaultPort)
	port := viper.GetInt(EnvPort)

	return port
}

func (apiConfiguration *APIConfiguration) validate() error {
	validator := validator.New()
	err := validator.Struct(apiConfiguration)

	return err
}
