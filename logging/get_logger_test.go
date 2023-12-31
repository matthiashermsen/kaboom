package logging_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	loggingConfiguration "github.com/matthiashermsen/kaboom/environment/logging"
	"github.com/matthiashermsen/kaboom/logging"
)

func TestGetLogger(suite *testing.T) {
	logLevels := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}

	for _, expectedLogLevel := range logLevels {
		suite.Run(fmt.Sprintf("Log level=%s", expectedLogLevel), func(testing *testing.T) {
			configuration := loggingConfiguration.LoggingConfiguration{
				Level: expectedLogLevel,
			}

			logger := logging.GetLogger(configuration)

			assert.NotNil(testing, logger, "Expected logger not to be nil")
		})
	}
}
