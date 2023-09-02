package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/kaboom/api/response"
)

func RequireJsonContentType(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			contentTypeKey := "Content-Type"
			jsonContentType := "application/json"

			actualContentType := request.Header.Get(contentTypeKey)

			if actualContentType != jsonContentType {
				responseWriter.WriteHeader(http.StatusUnsupportedMediaType)

				apiResponse := response.NewFailureApiResponse(response.ContentTypeInvalid, fmt.Sprintf("Expected '%s' to be '%s' but got '%s'", contentTypeKey, jsonContentType, actualContentType))
				err := response.WriteJsonResponse(responseWriter, apiResponse)

				if err != nil {
					response.HandleJsonResponseWriteError(responseWriter, err, logger)
				}

				return
			}

			next.ServeHTTP(responseWriter, request)
		})
	}
}
