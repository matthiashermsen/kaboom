package environment_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/environment"
)

func TestNewConfiguration(suite *testing.T) {
	suite.Run("Configuration with default values", func(testing *testing.T) {
		_, err := environment.NewConfiguration()

		assert.NoError(testing, err, "Expected no error when creating configuration")
	})
}
