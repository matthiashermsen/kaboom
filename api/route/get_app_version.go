package route

import "net/http"

func GetAppVersion(appVersion string) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")

		if appVersion == "" {
			// TODO OperationFailureResponse

			return
		}

		// TODO OperationSuccessResponse
	}
}
