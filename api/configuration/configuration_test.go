package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(suite *testing.T) {
	suite.Run("No error", func(testing *testing.T) {
		configuration := Configuration{Port: 3000}

		err := configuration.Validate()

		assert.NoError(testing, err, "Expected no validation error")
	})

	suite.Run("Invalid port", func(testing *testing.T) {
		configuration := Configuration{Port: 0}

		err := configuration.Validate()

		assert.Error(testing, err, "Expected validation error")
	})
}
