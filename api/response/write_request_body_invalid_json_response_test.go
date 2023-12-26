package response_test

import (
	"fmt"
	"log/slog"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/response"
)

func TestWriteRequestBodyInvalidJSONResponse(testing *testing.T) {
	responseWriter := httptest.NewRecorder()

	err := response.WriteRequestBodyInvalidJSONResponse(responseWriter, slog.New(slog.Default().Handler()))

	assert.Nil(testing, err, fmt.Sprintf("Expected error to be nil but got '%v'", err))

	expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"REQUEST_BODY_INVALID","message":"Invalid request body."}}`
	actualResponseBodyAsString := responseWriter.Body.String()

	assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
}
