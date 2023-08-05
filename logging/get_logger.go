package logging

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/matthiashermsen/kaboom/configuration/logging"
)

func GetLogger(configuration logging.Configuration) zerolog.Logger {
	return zerolog.New(os.Stderr).With().Timestamp().Logger().Level(configuration.Level)
}
