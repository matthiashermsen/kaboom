package ping

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func Handle(logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		apiResponse := response.NewSuccessAPIResponse("")
		response.WriteJSONResponse(responseWriter, apiResponse, logger)
	}
}
