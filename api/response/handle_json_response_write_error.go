package response

import (
	"log/slog"
	"net/http"
)

func HandleJsonResponseWriteError(responseWriter http.ResponseWriter, jsonResponseWriteError error, logger *slog.Logger) {
	logger.Error("Unable to write response", jsonResponseWriteError)

	apiResponse := NewErrorApiResponse()
	err := WriteJsonResponse(responseWriter, apiResponse)

	if err != nil {
		logger.Error("Unable to write error response", err)
	}
}
