package logging

import "github.com/rs/zerolog"

func LogError(logger zerolog.Logger, err error) {
	logger.Error().Err(err).Msg("")
}
