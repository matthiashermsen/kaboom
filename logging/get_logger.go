package logging

import (
	"log/slog"
	"os"

	"github.com/matthiashermsen/kaboom/configuration/logging"
)

func GetLogger(configuration logging.Configuration) *slog.Logger {
	textHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: configuration.Level})

	return slog.New(textHandler)
}
