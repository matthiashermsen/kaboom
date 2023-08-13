package api

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/matthiashermsen/kaboom/api/route"
)

func GetApi(appVersion string, logger *slog.Logger) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/app-version", route.GetAppVersion(appVersion, logger))

	return apiRouter
}
