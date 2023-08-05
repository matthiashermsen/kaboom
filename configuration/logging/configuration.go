package logging

import "github.com/rs/zerolog"

type Configuration struct {
	Level zerolog.Level `mapstructure:"LOGGING_LEVEL"`
}
