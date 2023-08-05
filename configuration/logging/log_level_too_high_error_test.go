package logging

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogLevelTooHighError_Error(testing *testing.T) {
	logLevelTooHighError := LogLevelTooHighError{}

	expectedErrorMessage := fmt.Sprintf("Log level is too high. Maximum log level is %d.", zerolog.PanicLevel)

	assert.Equal(testing, expectedErrorMessage, logLevelTooHighError.Error())
}
