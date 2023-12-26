package response

import (
	"log/slog"
	"net/http"
)

func WriteRequestBodyInvalidJSONResponse(responseWriter http.ResponseWriter, logger *slog.Logger) error {
	apiResponse := NewFailureAPIResponse(RequestBodyInvalid, "Invalid request body.")

	return WriteJSONResponse(responseWriter, apiResponse, logger)
}
