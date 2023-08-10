package response

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetHeaderContentTypeToJson(testing *testing.T) {
	responseWriter := httptest.NewRecorder()

	SetHeaderContentTypeToJson(responseWriter)

	expectedContentType := "application/json"
	actualContentType := responseWriter.Header().Get("Content-Type")

	assert.Equal(testing, expectedContentType, actualContentType, fmt.Sprintf("Expected Content-Type '%s', but got '%s'", expectedContentType, actualContentType))
}
