package response_test

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/response"
)

func TestWriteJSONResponse(suite *testing.T) {
	suite.Run("Success", func(testing *testing.T) {
		responseWriter := httptest.NewRecorder()
		responseBody := response.NewSuccessAPIResponse("made-up")

		err := response.WriteJSONResponse(responseWriter, responseBody, slog.New(slog.Default().Handler()))

		assert.Nil(testing, err, fmt.Sprintf("Expected error to be nil but got '%v'", err))

		expectedResponseBody := `{"status":"success","data":"made-up","error":null}`
		actualResponseBodyAsString := responseWriter.Body.String()

		assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
	})

	suite.Run("Marshalling failed", func(testing *testing.T) {
		responseWriter := httptest.NewRecorder()
		responseBody := response.NewSuccessAPIResponse(make(chan int))

		err := response.WriteJSONResponse(responseWriter, responseBody, slog.New(slog.Default().Handler()))

		assert.NotNil(testing, err, "Expected error not to be nil")
	})

	suite.Run("Write failed", func(testing *testing.T) {
		responseWriter := errorMockResponseWriter{}
		responseBody := response.NewSuccessAPIResponse("made-up")

		err := response.WriteJSONResponse(&responseWriter, responseBody, slog.New(slog.Default().Handler()))

		assert.NotNil(testing, err, "Expected error not to be nil")
	})
}

type errorMockResponseWriter struct {
	Body io.ReadCloser
	Code int
}

func (responseWriter *errorMockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (responseWriter *errorMockResponseWriter) Write(data []byte) (int, error) {
	responseWriter.Body = io.NopCloser(bytes.NewReader(make([]byte, 0)))
	return 0, fmt.Errorf("always errors")
}

func (responseWriter *errorMockResponseWriter) WriteHeader(statusCode int) {
	responseWriter.Code = statusCode
}
