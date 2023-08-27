package route

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func GetAppVersion(appVersion string, logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if appVersion == "" {
			apiResponse := response.NewFailureApiResponse(response.AppVersionUnavailable, "The app version is unavailable.")
			err := response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				response.HandleJsonResponseWriteError(responseWriter, err, logger)
			}

			return
		}

		apiResponse := response.NewSuccessApiResponse(appVersion)
		err := response.WriteJsonResponse(responseWriter, apiResponse)

		if err != nil {
			response.HandleJsonResponseWriteError(responseWriter, err, logger)
		}
	}
}
