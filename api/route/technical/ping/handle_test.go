package ping_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/route/technical/ping"
)

func TestHandleGetPing(suite *testing.T) {
	suite.Run("Ping", func(testing *testing.T) {
		logger := slog.New(slog.Default().Handler())

		request, err := http.NewRequest("GET", "/", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()
		handler := ping.Handle(logger)

		handler(responseRecorder, request)

		assert.Equal(testing, http.StatusOK, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusOK, responseRecorder.Code))

		expectedResponseBody := `{"status":"success","data":"","error":null}`
		actualResponseBody := responseRecorder.Body.String()

		assert.JSONEq(testing, expectedResponseBody, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBody))
	})
}
