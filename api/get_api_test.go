package api

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetApi(suite *testing.T) {
	suite.Run("Router is not nil", func(testing *testing.T) {
		router := GetApi("made-up", zerolog.Nop())

		assert.NotNil(testing, router, "Expected a non-nil router")
	})
}
