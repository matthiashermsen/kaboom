package response

import "net/http"

func SetHeaderContentTypeToJson(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}
