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
)

func main() {
	configuration, err := environment.NewConfiguration()

	if err != nil {
		panic(err)
	}

	logger := logging.GetLogger(configuration.Logging)

	apiRouter := api.GetAPI(appversion.AppVersion, logger)

	logger.Info(fmt.Sprintf("Starting server on port %v", configuration.API.Port))

	err = http.ListenAndServe(":"+strconv.Itoa(configuration.API.Port), apiRouter)

	if err != nil {
		logger.Error("", err)

		os.Exit(1)
	}
}
