package getappversion

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func Handle(appVersion string, logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if appVersion == "" {
			apiResponse := response.NewFailureAPIResponse(response.AppVersionUnavailable, "The app version is unavailable.")
			response.WriteJSONResponse(responseWriter, apiResponse, logger)

			return
		}

		apiResponse := response.NewSuccessAPIResponse(appVersion)
		response.WriteJSONResponse(responseWriter, apiResponse, logger)
	}
}
