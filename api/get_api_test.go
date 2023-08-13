package api

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApi(suite *testing.T) {
	suite.Run("Router is not nil", func(testing *testing.T) {
		router := GetApi("made-up", slog.New(slog.Default().Handler()))

		assert.NotNil(testing, router, "Expected a non-nil router")
	})
}
