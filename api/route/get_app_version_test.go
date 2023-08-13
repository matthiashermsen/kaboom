package route

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetAppVersion(suite *testing.T) {
	suite.Run("App version is available", func(testing *testing.T) {
		appVersion := "1.0.0"
		logger := zerolog.Nop()

		request, err := http.NewRequest("GET", "/", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()
		handler := GetAppVersion(appVersion, logger)

		handler(responseRecorder, request)

		assert.Equal(testing, http.StatusOK, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusOK, responseRecorder.Code))

		expectedContentType := "application/json"
		actualContentType := responseRecorder.Header().Get("Content-Type")

		assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))

		expectedResponseBody := `{"status":"success","data":"1.0.0","error":{"code":"","message":""}}`
		actualResponseBody := responseRecorder.Body.String()

		assert.JSONEq(testing, expectedResponseBody, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBody))
	})

	suite.Run("App version is unavailable", func(testing *testing.T) {
		appVersion := ""
		logger := zerolog.Nop()

		request, err := http.NewRequest("GET", "/", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()
		handler := GetAppVersion(appVersion, logger)

		handler(responseRecorder, request)

		assert.Equal(testing, http.StatusOK, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusOK, responseRecorder.Code))

		expectedContentType := "application/json"
		actualContentType := responseRecorder.Header().Get("Content-Type")

		assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))

		expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"APP_VERSION_UNAVAILABLE","message":"The app version is unavailable."}}`
		actualResponseBody := responseRecorder.Body.String()

		assert.JSONEq(testing, expectedResponseBody, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBody))
	})
}