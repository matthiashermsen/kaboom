package response

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse[T any](responseWriter http.ResponseWriter, responseBody ApiResponse[T]) error {
	encodedResponseBody, err := json.Marshal(responseBody)

	if err != nil {
		return err
	}

	_, err = responseWriter.Write(encodedResponseBody)

	return err
}
