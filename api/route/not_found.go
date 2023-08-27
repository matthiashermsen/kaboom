package route

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func RespondWithNotFound(logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusNotFound)

		apiResponse := response.NewFailureApiResponse(response.NotFound, fmt.Sprintf("Could not find '%s'", request.URL.Path))
		err := response.WriteJsonResponse(responseWriter, apiResponse)

		if err != nil {
			response.HandleJsonResponseWriteError(responseWriter, err, logger)
		}
	}
}
