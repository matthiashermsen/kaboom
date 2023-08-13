package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJsonResponse(suite *testing.T) {
	suite.Run("Success", func(testing *testing.T) {
		responseWriter := httptest.NewRecorder()
		responseBody := NewSuccessApiResponse("made-up")

		err := WriteJsonResponse(responseWriter, responseBody)

		assert.Nil(testing, err, fmt.Sprintf("Expected error to be nil but got '%v'", err))

		expectedResponseBody := `{"status":"success","data":"made-up","error":{"code":"","message":""}}`
		actualResponseBodyAsString := responseWriter.Body.String()

		assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
	})

	suite.Run("Marshalling failed", func(testing *testing.T) {
		responseWriter := httptest.NewRecorder()
		responseBody := NewSuccessApiResponse(make(chan int))

		err := WriteJsonResponse(responseWriter, responseBody)

		assert.NotNil(testing, err, "Expected error not to be nil")
	})

	suite.Run("Write failed", func(testing *testing.T) {
		responseWriter := errorMockResponseWriter{}
		responseBody := NewSuccessApiResponse("made-up")

		err := WriteJsonResponse(&responseWriter, responseBody)

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
