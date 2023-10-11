package api_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api"
)

func TestGetAPI(suite *testing.T) {
	suite.Run("Router is not nil", func(testing *testing.T) {
		router := api.GetAPI("made-up", slog.New(slog.Default().Handler()))

		assert.NotNil(testing, router, "Expected a non-nil router")
	})

	suite.Run("Technical endpoints", func(technicalEndpointSuite *testing.T) {
		technicalEndpointSuite.Run("GetPing", func(getPingSuite *testing.T) {
			getPingSuite.Run("Responds with Content-Type application/json", func(testing *testing.T) {
				router := api.GetAPI("made-up", slog.New(slog.Default().Handler()))
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
				router := api.GetAPI("made-up", slog.New(slog.Default().Handler()))
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
				router := api.GetAPI("made-up", slog.New(slog.Default().Handler()))
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
