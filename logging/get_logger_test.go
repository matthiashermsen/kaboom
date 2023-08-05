package logging

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/configuration/logging"
)

func TestGetLogger_LogLevels(testingSuite *testing.T) {
	logLevels := []zerolog.Level{
		zerolog.PanicLevel,
		zerolog.FatalLevel,
		zerolog.ErrorLevel,
		zerolog.WarnLevel,
		zerolog.InfoLevel,
		zerolog.DebugLevel,
		zerolog.TraceLevel,
	}

	for _, expectedLogLevel := range logLevels {
		testingSuite.Run(fmt.Sprintf("Log level=%s", expectedLogLevel), func(testing *testing.T) {
			configuration := logging.Configuration{
				Level: expectedLogLevel,
			}

			logger := GetLogger(configuration)
			actualLogLevel := logger.GetLevel()

			assert.Equal(testing, expectedLogLevel, actualLogLevel, fmt.Sprintf("Expected logger to have log level %s but got %s", expectedLogLevel, actualLogLevel))
		})
	}
}
