package command

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureNewBoard(suite *testing.T) {
	suite.Run("TODO", func(testing *testing.T) {
		logger := slog.New(slog.Default().Handler())

		request, err := http.NewRequest("GET", "/", nil)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()
		handler := HandleConfigureNewBoard(logger)

		handler(responseRecorder, request)

		assert.Equal(testing, http.StatusNotImplemented, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusNotImplemented, responseRecorder.Code))
	})
}
