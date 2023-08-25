package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApi(suite *testing.T) {
	suite.Run("Router is not nil", func(testing *testing.T) {
		router := GetApi("made-up", slog.New(slog.Default().Handler()))

		assert.NotNil(testing, router, "Expected a non-nil router")
	})
}

func TestGetPing(suite *testing.T) {
	suite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
		router := GetApi("made-up", slog.New(slog.Default().Handler()))
		request, err := http.NewRequest("GET", "/ping", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		expectedContentType := "application/json"
		actualContentType := responseRecorder.Header().Get("Content-Type")

		assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
	})
}

func TestGetAppVersion(suite *testing.T) {
	suite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
		router := GetApi("made-up", slog.New(slog.Default().Handler()))
		request, err := http.NewRequest("GET", "/app-version", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		expectedContentType := "application/json"
		actualContentType := responseRecorder.Header().Get("Content-Type")

		assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
	})
}
