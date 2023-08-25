package api

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/matthiashermsen/kaboom/api/middleware"
	"github.com/matthiashermsen/kaboom/api/route"
)

func GetApi(appVersion string, logger *slog.Logger) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Use(middleware.SetJSONContentType)

	apiRouter.Get("/ping", route.GetPing(logger))
	apiRouter.Get("/app-version", route.GetAppVersion(appVersion, logger))

	return apiRouter
}
