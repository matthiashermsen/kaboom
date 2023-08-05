package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortInvalidError_Error(testing *testing.T) {
	portInvalidError := PortInvalidError{}

	expectedErrorMessage := fmt.Sprintf("Invalid port. Port number should be between %d and %d (inclusive).", minimumPort, maximumPort)

	assert.Equal(testing, expectedErrorMessage, portInvalidError.Error())
}
