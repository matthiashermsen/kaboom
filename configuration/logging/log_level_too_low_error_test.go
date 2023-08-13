package logging

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogLevelTooLowError(testing *testing.T) {
	logLevelTooLowError := LogLevelTooLowError{}

	expectedErrorMessage := fmt.Sprintf("Log level is too low. Minimum log level is %d.", zerolog.TraceLevel)

	assert.Equal(testing, expectedErrorMessage, logLevelTooLowError.Error())
}
