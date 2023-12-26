package startnewgame_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/kaboom/api/route/command/startnewgame"
	"github.com/matthiashermsen/kaboom/storage/inmemory"
)

func TestHandle(suite *testing.T) {
	suite.Run("SAMPLE TEST", func(testing *testing.T) {
		logger := slog.New(slog.Default().Handler())
		store := inmemory.New()

		requestBody := startnewgame.RequestBody{
			AmountOfRows:    2,
			AmountOfColumns: 2,
			AmountOfMines:   2,
		}
		encodedRequestBody, err := json.Marshal(requestBody)

		assert.NoError(testing, err, "Expected no error when marshaling request body")

		encodedRequestBodyBytes := bytes.NewBuffer(encodedRequestBody)
		request, err := http.NewRequest("POST", "/start-new-game", encodedRequestBodyBytes)

		assert.NoError(testing, err, "Expected no error when constructing request")

		responseRecorder := httptest.NewRecorder()
		handler := startnewgame.Handle(store, logger)

		handler(responseRecorder, request)

		assert.Equal(testing, http.StatusOK, responseRecorder.Code, fmt.Sprintf("Expected status code %d but got %d", http.StatusOK, responseRecorder.Code))

		// expectedResponseBody := `{"status":"success","data":"Ok","error":null}`
		// actualResponseBody := responseRecorder.Body.String()

		// assert.JSONEq(testing, expectedResponseBody, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedResponseBody, actualResponseBody))
	})
}
