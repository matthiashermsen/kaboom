package logging

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestValidateConfiguration(suite *testing.T) {
	suite.Run("No error", func(testing *testing.T) {
		configuration := Configuration{Level: zerolog.WarnLevel}

		err := ValidateConfiguration(configuration)

		assert.NoError(testing, err, "Expected no validation error")
	})

	suite.Run("Log level too low", func(suite *testing.T) {
		invalidLogLevels := []zerolog.Level{
			zerolog.Level(-2),
			zerolog.Level(-5),
			zerolog.Level(-10),
		}

		for _, currentLogLevel := range invalidLogLevels {
			suite.Run(fmt.Sprintf("Log level=%s", currentLogLevel), func(testing *testing.T) {
				configuration := Configuration{Level: currentLogLevel}

				err := ValidateConfiguration(configuration)

				assert.Error(testing, err, "Expected validation error")
				assert.IsType(testing, LogLevelTooLowError{}, err, "Expected LogLevelTooLowError")
			})
		}
	})

	suite.Run("Log level too high", func(suite *testing.T) {
		invalidLogLevels := []zerolog.Level{
			zerolog.Level(6),
			zerolog.Level(10),
			zerolog.Level(100),
		}

		for _, currentLogLevel := range invalidLogLevels {
			suite.Run(fmt.Sprintf("Log level=%s", currentLogLevel), func(testing *testing.T) {
				configuration := Configuration{Level: currentLogLevel}

				err := ValidateConfiguration(configuration)

				assert.Error(testing, err, "Expected validation error")
				assert.IsType(testing, LogLevelTooHighError{}, err, "Expected LogLevelTooHighError")
			})
		}
	})
}
