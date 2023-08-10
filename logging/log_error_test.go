package logging

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogError(testing *testing.T) {
	logger := zerolog.New(os.Stdout)
	testError := errors.New("test error")
	logHook := &logHook{}

	logger = logger.Hook(logHook)

	LogError(logger, testError)

	amountOfLogEvents := len(logHook.logEvents)

	assert.True(testing, amountOfLogEvents == 1, fmt.Sprintf("Expected to have exactly one log event but got %d events", amountOfLogEvents))
}

type logHook struct {
	logEvents []zerolog.Event
}

func (logHook *logHook) Run(logEvent *zerolog.Event, level zerolog.Level, message string) {
	logHook.logEvents = append(logHook.logEvents, *logEvent)
}
