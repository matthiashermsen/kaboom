package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	"github.com/matthiashermsen/kaboom/api/route"
)

func GetApi(appVersion string, logger zerolog.Logger) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/app-version", route.GetAppVersion(appVersion, logger))

	return apiRouter
}
