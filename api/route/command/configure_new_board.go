package command

import (
	"log/slog"
	"net/http"
)

func HandleConfigureNewBoard(logger *slog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusNotImplemented)
	}
}
