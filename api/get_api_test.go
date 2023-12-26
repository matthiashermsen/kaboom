package api_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api"
	"github.com/matthiashermsen/kaboom/storage/inmemory"
)

func TestGetAPI(suite *testing.T) {
	suite.Run("Router is not nil", func(testing *testing.T) {
		router := api.GetAPI(inmemory.New(), slog.New(slog.Default().Handler()), "made-up")

		assert.NotNil(testing, router, "Expected a non-nil router")
	})

	suite.Run("Command endpoints", func(commandEndpointSuite *testing.T) {
		commandEndpointSuite.Run("StartNewGame", func(startNewGameSuite *testing.T) {
			startNewGameSuite.Run("Requires Content-Type application/json", func(testing *testing.T) {
				router := api.GetAPI(inmemory.New(), slog.New(slog.Default().Handler()), "made-up")

				request, err := http.NewRequest("POST", "/command/start-new-game", strings.NewReader("made-up"))

				assert.NoError(testing, err, "Expected no error when constructing request")

				responseRecorder := httptest.NewRecorder()

				router.ServeHTTP(responseRecorder, request)

				assert.Equal(testing, http.StatusUnsupportedMediaType, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusUnsupportedMediaType, responseRecorder.Code))

				expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"CONTENT_TYPE_INVALID","message":"Expected 'Content-Type' to be 'application/json' but got ''"}}`
				actualResponseBodyAsString := responseRecorder.Body.String()

				assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
			})

			startNewGameSuite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
				router := api.GetAPI(inmemory.New(), slog.New(slog.Default().Handler()), "made-up")

				request, err := http.NewRequest("POST", "/command/start-new-game", strings.NewReader("made-up"))

				assert.NoError(testing, err, "Expected no error when constructing request")

				responseRecorder := httptest.NewRecorder()

				router.ServeHTTP(responseRecorder, request)

				expectedContentType := "application/json"
				actualContentType := responseRecorder.Header().Get("Content-Type")

				assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
			})
		})
	})

	suite.Run("Technical endpoints", func(technicalEndpointSuite *testing.T) {
		technicalEndpointSuite.Run("GetPing", func(getPingSuite *testing.T) {
			getPingSuite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
				router := api.GetAPI(inmemory.New(), slog.New(slog.Default().Handler()), "made-up")
				request, err := http.NewRequest("GET", "/ping", nil)

				assert.NoError(testing, err, "Expected no error when constructing request")

				responseRecorder := httptest.NewRecorder()

				router.ServeHTTP(responseRecorder, request)

				expectedContentType := "application/json"
				actualContentType := responseRecorder.Header().Get("Content-Type")

				assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
			})
		})

		technicalEndpointSuite.Run("GetAppVersion", func(getAppVersionSuite *testing.T) {
			getAppVersionSuite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
				router := api.GetAPI(inmemory.New(), slog.New(slog.Default().Handler()), "made-up")
				request, err := http.NewRequest("GET", "/app-version", nil)

				assert.NoError(testing, err, "Expected no error when constructing request")

				responseRecorder := httptest.NewRecorder()

				router.ServeHTTP(responseRecorder, request)

				expectedContentType := "application/json"
				actualContentType := responseRecorder.Header().Get("Content-Type")

				assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
			})
		})

		technicalEndpointSuite.Run("RespondWithNotFound", func(respondWithNotFoundSuite *testing.T) {
			respondWithNotFoundSuite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
				router := api.GetAPI(inmemory.New(), slog.New(slog.Default().Handler()), "made-up")
				request, err := http.NewRequest("GET", "/made-up", nil)

				assert.NoError(testing, err, "Expected no error when constructing request")

				responseRecorder := httptest.NewRecorder()

				router.ServeHTTP(responseRecorder, request)

				expectedContentType := "application/json"
				actualContentType := responseRecorder.Header().Get("Content-Type")

				assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
			})
		})
	})
}
