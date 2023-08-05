package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateConfiguration_NoError(testing *testing.T) {
	configuration := Configuration{Port: 3000}

	err := ValidateConfiguration(configuration)

	assert.NoError(testing, err, "Expected no validation error")
}

func TestValidateConfiguration_InvalidPort(testing *testing.T) {
	configuration := Configuration{Port: 0}

	err := ValidateConfiguration(configuration)

	assert.Error(testing, err, "Expected validation error")
	assert.IsType(testing, PortInvalidError{}, err, "Expected PortInvalidError")
}
