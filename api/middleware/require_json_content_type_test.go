package middleware_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/middleware"
)

func TestRequireJSONContentType(suite *testing.T) {
	handler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusOK)
	})

	logger := slog.New(slog.Default().Handler())

	middleware := middleware.RequireJSONContentType(logger)

	suite.Run("'Content-Type' is 'application/json'", func(testing *testing.T) {
		request, err := http.NewRequest("GET", "/made-up", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		request.Header.Set("Content-Type", "application/json")

		responseRecorder := httptest.NewRecorder()

		middleware(handler).ServeHTTP(responseRecorder, request)

		assert.Equal(testing, http.StatusOK, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusOK, responseRecorder.Code))
	})

	suite.Run("'Content-Type' is not 'application/json'", func(testing *testing.T) {
		request, err := http.NewRequest("GET", "/made-up", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		request.Header.Set("Content-Type", "made-up")

		responseRecorder := httptest.NewRecorder()

		middleware(handler).ServeHTTP(responseRecorder, request)

		assert.Equal(testing, http.StatusUnsupportedMediaType, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusUnsupportedMediaType, responseRecorder.Code))

		expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"CONTENT_TYPE_INVALID","message":"Expected 'Content-Type' to be 'application/json' but got 'made-up'"}}`
		actualResponseBodyAsString := responseRecorder.Body.String()

		assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
	})

	suite.Run("'Content-Type' is empty", func(testing *testing.T) {
		request, err := http.NewRequest("GET", "/made-up", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		request.Header.Set("Content-Type", "")

		responseRecorder := httptest.NewRecorder()

		middleware(handler).ServeHTTP(responseRecorder, request)

		assert.Equal(testing, http.StatusUnsupportedMediaType, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusUnsupportedMediaType, responseRecorder.Code))

		expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"CONTENT_TYPE_INVALID","message":"Expected 'Content-Type' to be 'application/json' but got ''"}}`
		actualResponseBodyAsString := responseRecorder.Body.String()

		assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
	})

	suite.Run("'Content-Type' is missing", func(testing *testing.T) {
		request, err := http.NewRequest("GET", "/made-up", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()

		middleware(handler).ServeHTTP(responseRecorder, request)

		assert.Equal(testing, http.StatusUnsupportedMediaType, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusUnsupportedMediaType, responseRecorder.Code))

		expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"CONTENT_TYPE_INVALID","message":"Expected 'Content-Type' to be 'application/json' but got ''"}}`
		actualResponseBodyAsString := responseRecorder.Body.String()

		assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
	})
}
