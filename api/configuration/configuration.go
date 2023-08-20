package configuration

import "github.com/go-playground/validator/v10"

type Configuration struct {
	Port uint16 `mapstructure:"SERVER_PORT" validate:"gte=1,lte=65535"`
}

func (configuraton *Configuration) Validate() error {
	validate := validator.New()

	return validate.Struct(configuraton)
}
