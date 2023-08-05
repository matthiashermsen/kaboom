package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/matthiashermsen/kaboom/api/route"
)

func GetApi(appVersion string) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/app-version", route.GetAppVersion(appVersion))

	return apiRouter
}
