package route

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func GetPing(logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		apiResponse := response.NewSuccessApiResponse("")
		err := response.WriteJsonResponse(responseWriter, apiResponse)

		if err != nil {
			response.HandleJsonResponseWriteError(responseWriter, err, logger)
		}
	}
}
