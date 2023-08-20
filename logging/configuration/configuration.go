package configuration

import (
	"log/slog"

	"github.com/go-playground/validator/v10"

	"github.com/matthiashermsen/kaboom/logging/validation"
)

type Configuration struct {
	Level slog.Level `mapstructure:"LOGGING_LEVEL" validate:"logLevel"`
}

func (configuraton *Configuration) Validate() error {
	validate := validator.New()

	err := validate.RegisterValidation("logLevel", func(fieldLevel validator.FieldLevel) bool {
		configuredLogLevel := fieldLevel.Field().Interface()

		return validation.IsLogLevel(configuredLogLevel)
	})

	if err != nil {
		return err
	}

	return validate.Struct(configuraton)
}
