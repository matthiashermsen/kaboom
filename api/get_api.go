package api

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/matthiashermsen/kaboom/api/middleware"
	"github.com/matthiashermsen/kaboom/api/route/command/startnewgame"
	"github.com/matthiashermsen/kaboom/api/route/technical/getappversion"
	"github.com/matthiashermsen/kaboom/api/route/technical/notfound"
	"github.com/matthiashermsen/kaboom/api/route/technical/ping"
	"github.com/matthiashermsen/kaboom/storage"
)

func GetAPI(store storage.Store, logger *slog.Logger, appVersion string) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Use(middleware.SetJSONContentType)

	apiRouter.Route("/command", func(commandRouter chi.Router) {
		commandRouter.Use(middleware.RequireJSONContentType(logger))

		commandRouter.Post("/start-new-game", startnewgame.Handle(store, logger))
	})

	apiRouter.Get("/ping", ping.Handle(logger))
	apiRouter.Get("/app-version", getappversion.Handle(appVersion, logger))

	apiRouter.NotFound(notfound.Handle(logger))

	return apiRouter
}
