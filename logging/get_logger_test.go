package logging

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/configuration/logging"
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
			configuration := logging.Configuration{
				Level: expectedLogLevel,
			}

			logger := GetLogger(configuration)

			assert.NotNil(testing, logger, "Expected logger not to be nil")
		})
	}
}
