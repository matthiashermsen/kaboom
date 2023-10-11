package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func RequireJSONContentType(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			contentTypeKey := "Content-Type"
			jsonContentType := "application/json"

			actualContentType := request.Header.Get(contentTypeKey)

			if actualContentType != jsonContentType {
				responseWriter.WriteHeader(http.StatusUnsupportedMediaType)

				apiResponse := response.NewFailureAPIResponse(response.ContentTypeInvalid, fmt.Sprintf("Expected '%s' to be '%s' but got '%s'", contentTypeKey, jsonContentType, actualContentType))
				response.WriteJSONResponse(responseWriter, apiResponse, logger)

				return
			}

			next.ServeHTTP(responseWriter, request)
		})
	}
}
