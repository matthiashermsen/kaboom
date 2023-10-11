package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteJSONResponse[T any](responseWriter http.ResponseWriter, responseBody APIResponse[T], logger *slog.Logger) error {
	encodedResponseBody, err := json.Marshal(responseBody)

	if err != nil {
		return err
	}

	_, err = responseWriter.Write(encodedResponseBody)

	if err == nil {
		return nil
	}

	logger.Error("Unable to write response", err)

	errorResponseBody := NewErrorAPIResponse()
	encodedErrorResponseBody, err := json.Marshal(errorResponseBody)

	if err != nil {
		return err
	}

	_, err = responseWriter.Write(encodedErrorResponseBody)

	return err
}
