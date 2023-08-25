package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetJSONContentType(testing *testing.T) {
	handler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusOK)
	})

	middlewareHandler := SetJSONContentType(handler)

	request := httptest.NewRequest("GET", "/", nil)
	responseWriter := httptest.NewRecorder()

	middlewareHandler.ServeHTTP(responseWriter, request)

	response := responseWriter.Result()

	expectedContentType := "application/json"
	actualContentType := response.Header.Get("Content-Type")

	assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
}
