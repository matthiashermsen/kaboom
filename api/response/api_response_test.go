package response

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccessApiResponse(testing *testing.T) {
	data := "made-up"

	apiResponse := NewSuccessApiResponse(data)

	assert.Equal(testing, "success", apiResponse.Status, fmt.Sprintf("Expected status to be 'success' but got '%s'", apiResponse.Status))
	assert.Equal(testing, data, apiResponse.Data, fmt.Sprintf("Expected data to be '%s' but got '%s'", data, apiResponse.Data))
	assert.Equal(testing, "", apiResponse.Error.Code, "Expected error code to be empty")
	assert.Equal(testing, "", apiResponse.Error.Message, "Expected error message to be empty")
}

func TestNewFailureApiResponse(testing *testing.T) {
	errorCode := "SOMETHING_FAILED"
	errorMessage := "error message"

	apiResponse := NewFailureApiResponse(errorCode, errorMessage)

	assert.Equal(testing, "failure", apiResponse.Status, fmt.Sprintf("Expected status to be 'failure' but got '%s'", apiResponse.Status))
	assert.Nil(testing, apiResponse.Data, "Expected data to be nil")
	assert.Equal(testing, errorCode, apiResponse.Error.Code, fmt.Sprintf("Expected error code to be '%s' but got '%s'", errorCode, apiResponse.Error.Code))
	assert.Equal(testing, errorMessage, apiResponse.Error.Message, fmt.Sprintf("Expected error message to be '%s' but got '%s'", errorMessage, apiResponse.Error.Message))
}

func TestNewErrorApiResponse(testing *testing.T) {
	apiResponse := NewErrorApiResponse()

	assert.Equal(testing, "error", apiResponse.Status, fmt.Sprintf("Expected status to be 'error' but got '%s'", apiResponse.Status))
	assert.Nil(testing, apiResponse.Data, "Expected data to be nil")

	expectedErrorMessage := "The server encountered an unexpected condition that prevented it from fulfilling the request."

	assert.Equal(testing, InternalError, apiResponse.Error.Code, fmt.Sprintf("Expected error code to be '%s' but got '%s'", InternalError, apiResponse.Error.Code))
	assert.Equal(testing, expectedErrorMessage, apiResponse.Error.Message, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrorMessage, apiResponse.Error.Message))
}
