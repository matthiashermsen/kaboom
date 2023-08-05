package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApi(suite *testing.T) {
	suite.Run("Router is not nil", func(testing *testing.T) {
		router := GetApi("made-up")

		assert.NotNil(testing, router, "Expected a non-nil router")
	})
}
