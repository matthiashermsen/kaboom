package route

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/matthiashermsen/kaboom/api/response"
	"github.com/matthiashermsen/kaboom/logging"
)

func GetAppVersion(appVersion string, logger zerolog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		response.SetHeaderContentTypeToJson(responseWriter)

		if appVersion == "" {
			apiResponse := response.NewFailureApiResponse("APP_VERSION_UNAVAILABLE", "The app version is unavailable.")
			err := response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				logging.LogError(logger, err)

				apiResponse = response.NewErrorApiResponse()
				err = response.WriteJsonResponse(responseWriter, apiResponse)

				if err != nil {
					logging.LogError(logger, err)
				}
			}

			return
		}

		apiResponse := response.NewSuccessApiResponse(appVersion)
		err := response.WriteJsonResponse(responseWriter, apiResponse)

		if err != nil {
			logging.LogError(logger, err)

			apiResponse := response.NewErrorApiResponse()
			err = response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				logging.LogError(logger, err)
			}
		}
	}
}
