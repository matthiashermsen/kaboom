package logging

import (
	"log/slog"
	"os"

	"github.com/matthiashermsen/kaboom/logging/configuration"
)

func GetLogger(configuration configuration.Configuration) *slog.Logger {
	textHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: configuration.Level})

	return slog.New(textHandler)
}
