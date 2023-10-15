package notfound_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/route/technical/notfound"
)

func TestHandle(suite *testing.T) {
	suite.Run("Respond with not found", func(testing *testing.T) {
		logger := slog.New(slog.Default().Handler())

		request, err := http.NewRequest("GET", "/made-up", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()
		handler := notfound.Handle(logger)

		handler(responseRecorder, request)

		assert.Equal(testing, http.StatusNotFound, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusNotFound, responseRecorder.Code))

		expectedResponseBody := `{"status":"failure","data":null,"error":{"code":"NOT_FOUND","message":"Could not find '/made-up'"}}`
		actualResponseBody := responseRecorder.Body.String()

		assert.JSONEq(testing, expectedResponseBody, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBody))
	})
}
