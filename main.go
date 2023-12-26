package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/matthiashermsen/kaboom/api"
	"github.com/matthiashermsen/kaboom/appversion"
	"github.com/matthiashermsen/kaboom/environment"
	"github.com/matthiashermsen/kaboom/logging"
	"github.com/matthiashermsen/kaboom/storage/inmemory"
)

func main() {
	configuration, err := environment.NewConfiguration()

	if err != nil {
		panic(err)
	}

	logger := logging.GetLogger(configuration.Logging)

	store := inmemory.New()

	apiRouter := api.GetAPI(store, logger, appversion.AppVersion)

	logger.Info(fmt.Sprintf("Starting server on port %v", configuration.API.Port))

	err = http.ListenAndServe(":"+strconv.Itoa(configuration.API.Port), apiRouter)

	if err != nil {
		logger.Error("", err)

		os.Exit(1)
	}
}
