package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/matthiashermsen/kaboom/api"
	apiConfiguration "github.com/matthiashermsen/kaboom/api/configuration"
	"github.com/matthiashermsen/kaboom/appversion"
	"github.com/matthiashermsen/kaboom/configuration"
	"github.com/matthiashermsen/kaboom/logging"
	loggingConfiguration "github.com/matthiashermsen/kaboom/logging/configuration"
)

func main() {
	configuration.InitializeConfigurationEnvironment()

	loggingConfiguration, err := loggingConfiguration.GetConfiguration()

	if err != nil {
		panic(err)
	}

	err = loggingConfiguration.Validate()

	if err != nil {
		panic(err)
	}

	logger := logging.GetLogger(loggingConfiguration)

	apiConfiguration, err := apiConfiguration.GetConfiguration()

	if err != nil {
		logErrorAndExit(err, logger)
	}

	err = apiConfiguration.Validate()

	if err != nil {
		logErrorAndExit(err, logger)
	}

	apiRouter := api.GetApi(appversion.AppVersion, logger)

	logger.Info(fmt.Sprintf("Starting server on port %v", apiConfiguration.Port))

	err = http.ListenAndServe(":"+strconv.Itoa(int(apiConfiguration.Port)), apiRouter)

	if err != nil {
		logErrorAndExit(err, logger)
	}
}

func logErrorAndExit(err error, logger *slog.Logger) {
	logger.Error("", err)

	os.Exit(1)
}
