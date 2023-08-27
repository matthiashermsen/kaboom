package response

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleJsonResponseWriteError(suite *testing.T) {
	suite.Run("Sends error response", func(testing *testing.T) {
		responseWriter := httptest.NewRecorder()
		logger := slog.New(slog.Default().Handler())
		err := errors.New("made-up")

		HandleJsonResponseWriteError(responseWriter, err, logger)

		expectedResponseBody := `{"status":"error","data":null,"error":{"code":"INTERNAL_ERROR","message":"The server encountered an unexpected condition that prevented it from fulfilling the request."}}`
		actualResponseBodyAsString := responseWriter.Body.String()

		assert.Equal(testing, expectedResponseBody, actualResponseBodyAsString, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBodyAsString))
	})
}
