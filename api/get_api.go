package api

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/matthiashermsen/kaboom/api/middleware"
	"github.com/matthiashermsen/kaboom/api/route/command"
	"github.com/matthiashermsen/kaboom/api/route/technical"
)

func GetApi(appVersion string, logger *slog.Logger) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Use(middleware.SetJsonContentType)

	apiRouter.Route("/command", func(commandRouter chi.Router) {
		commandRouter.Use(middleware.RequireJsonContentType(logger))

		commandRouter.Post("/configure-new-board", command.HandleConfigureNewBoard(logger))
	})

	apiRouter.Get("/ping", technical.HandleGetPing(logger))
	apiRouter.Get("/app-version", technical.HandleGetAppVersion(appVersion, logger))

	apiRouter.NotFound(technical.HandleNotFound(logger))

	return apiRouter
}
