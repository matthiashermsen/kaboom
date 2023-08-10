package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog"

	"github.com/matthiashermsen/kaboom/api"
	"github.com/matthiashermsen/kaboom/appversion"
	"github.com/matthiashermsen/kaboom/configuration"
	loggingConfig "github.com/matthiashermsen/kaboom/configuration/logging"
	"github.com/matthiashermsen/kaboom/configuration/server"
	"github.com/matthiashermsen/kaboom/logging"
)

func main() {
	configuration.InitializeConfigurationEnvironment()

	loggingConfiguration, err := loggingConfig.GetConfiguration()

	if err != nil {
		panic(err)
	}

	err = loggingConfig.ValidateConfiguration(loggingConfiguration)

	if err != nil {
		panic(err)
	}

	logger := logging.GetLogger(loggingConfiguration)

	serverConfiguration, err := server.GetConfiguration()

	if err != nil {
		logErrorAndExit(err, logger)
	}

	err = server.ValidateConfiguration(serverConfiguration)

	if err != nil {
		logErrorAndExit(err, logger)
	}

	apiRouter := api.GetApi(appversion.AppVersion, logger)

	logger.Info().Msg(fmt.Sprintf("Starting server on port %v", serverConfiguration.Port))

	err = http.ListenAndServe(":"+strconv.Itoa(int(serverConfiguration.Port)), apiRouter)

	if err != nil {
		logErrorAndExit(err, logger)
	}
}

func logErrorAndExit(err error, logger zerolog.Logger) {
	logger.Error().Err(err).Msg("")

	os.Exit(1)
}
