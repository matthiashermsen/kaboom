package route

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppVersion(suite *testing.T) {
	suite.Run("Test with appVersion 1.0.0", func(testing *testing.T) {
		req, err := http.NewRequest("GET", "/app-version", nil)

		assert.NoError(testing, err, "Expected no error when creating request")

		requestRecorder := httptest.NewRecorder()
		requestHandler := GetAppVersion("1.0.0")

		requestHandler.ServeHTTP(requestRecorder, req)

		assert.Equal(testing, http.StatusOK, requestRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusOK, requestRecorder.Code))

		expectedContentType := "application/json"
		actualContentType := requestRecorder.Header().Get("Content-Type")

		assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type %s but got %s", expectedContentType, actualContentType))
	})
}
