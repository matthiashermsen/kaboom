package logging

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestValidateConfiguration_NoError(testing *testing.T) {
	configuration := Configuration{Level: zerolog.WarnLevel}

	err := ValidateConfiguration(configuration)

	assert.NoError(testing, err, "Expected no validation error")
}

func TestValidateConfiguration_LogLevelTooLow(testing *testing.T) {
	invalidLogLevels := []zerolog.Level{
		zerolog.Level(-2),
		zerolog.Level(-5),
		zerolog.Level(-10),
	}

	for _, currentLogLevel := range invalidLogLevels {
		configuration := Configuration{Level: currentLogLevel}

		err := ValidateConfiguration(configuration)

		assert.Error(testing, err, "Expected validation error")
		assert.IsType(testing, LogLevelTooLowError{}, err, "Expected LogLevelTooLowError")
	}
}

func TestValidateConfiguration_LogLevelTooHigh(testing *testing.T) {
	invalidLogLevels := []zerolog.Level{
		zerolog.Level(6),
		zerolog.Level(10),
		zerolog.Level(100),
	}

	for _, currentLogLevel := range invalidLogLevels {
		configuration := Configuration{Level: currentLogLevel}

		err := ValidateConfiguration(configuration)

		assert.Error(testing, err, "Expected validation error")
		assert.IsType(testing, LogLevelTooHighError{}, err, "Expected LogLevelTooHighError")
	}
}
