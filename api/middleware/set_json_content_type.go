package middleware

import "net/http"

func SetJSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(responseWriter, request)
	})
}
