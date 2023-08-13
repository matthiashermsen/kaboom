package route

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func GetAppVersion(appVersion string, logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		response.SetHeaderContentTypeToJson(responseWriter)

		if appVersion == "" {
			apiResponse := response.NewFailureApiResponse(response.AppVersionUnavailable, "The app version is unavailable.")
			err := response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				logger.Error("Unable to write failure response", err)

				apiResponse = response.NewErrorApiResponse()
				err = response.WriteJsonResponse(responseWriter, apiResponse)

				if err != nil {
					logger.Error("Unable to write error response", err)
				}
			}

			return
		}

		apiResponse := response.NewSuccessApiResponse(appVersion)
		err := response.WriteJsonResponse(responseWriter, apiResponse)

		if err != nil {
			logger.Error("Unable to write success response", err)

			apiResponse := response.NewErrorApiResponse()
			err = response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				logger.Error("Unable to write error response", err)
			}
		}
	}
}
