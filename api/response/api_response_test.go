package response_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/response"
)

func TestNewSuccessAPIResponse(testing *testing.T) {
	data := "made-up"

	apiResponse := response.NewSuccessAPIResponse(data)

	assert.Equal(testing, "success", apiResponse.Status, fmt.Sprintf("Expected status to be 'success' but got '%s'", apiResponse.Status))
	assert.Equal(testing, data, apiResponse.Data, fmt.Sprintf("Expected data to be '%s' but got '%s'", data, apiResponse.Data))
	assert.Nil(testing, apiResponse.Error, "Expected error to be nil")
}

func TestNewFailureAPIResponse(testing *testing.T) {
	errorCode := "SOMETHING_FAILED"
	errorMessage := "error message"

	apiResponse := response.NewFailureAPIResponse(errorCode, errorMessage)

	assert.Equal(testing, "failure", apiResponse.Status, fmt.Sprintf("Expected status to be 'failure' but got '%s'", apiResponse.Status))
	assert.Nil(testing, apiResponse.Data, "Expected data to be nil")
	assert.Equal(testing, errorCode, apiResponse.Error.Code, fmt.Sprintf("Expected error code to be '%s' but got '%s'", errorCode, apiResponse.Error.Code))
	assert.Equal(testing, errorMessage, apiResponse.Error.Message, fmt.Sprintf("Expected error message to be '%s' but got '%s'", errorMessage, apiResponse.Error.Message))
}

func TestNewErrorAPIResponse(testing *testing.T) {
	apiResponse := response.NewErrorAPIResponse()

	assert.Equal(testing, "error", apiResponse.Status, fmt.Sprintf("Expected status to be 'error' but got '%s'", apiResponse.Status))
	assert.Nil(testing, apiResponse.Data, "Expected data to be nil")

	expectedErrorMessage := "The server encountered an unexpected condition that prevented it from fulfilling the request."

	assert.Equal(testing, response.InternalError, apiResponse.Error.Code, fmt.Sprintf("Expected error code to be '%s' but got '%s'", response.InternalError, apiResponse.Error.Code))
	assert.Equal(testing, expectedErrorMessage, apiResponse.Error.Message, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrorMessage, apiResponse.Error.Message))
}
