package notfound

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func Handle(logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusNotFound)

		apiResponse := response.NewFailureAPIResponse(response.NotFound, fmt.Sprintf("Could not find '%s'", request.URL.Path))
		response.WriteJSONResponse(responseWriter, apiResponse, logger)
	}
}
