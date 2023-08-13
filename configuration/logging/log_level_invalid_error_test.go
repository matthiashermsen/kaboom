package logging

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogLevelInvalidError(testing *testing.T) {
	logLevelInvalidError := LogLevelInvalidError{}

	expectedErrorMessage := fmt.Sprintf("Log level is invalid. Must be one of [%d, %d, %d, %d].", slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError)

	assert.Equal(testing, expectedErrorMessage, logLevelInvalidError.Error())
}
