package logging

import (
	"log/slog"
	"os"

	"github.com/matthiashermsen/kaboom/environment/logging"
)

func GetLogger(configuration logging.LoggingConfiguration) *slog.Logger {
	textHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: configuration.Level})

	return slog.New(textHandler)
}
