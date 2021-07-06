package engine

import (
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointsIsValid(t *testing.T) {
	t.Run("valid points", func(t *testing.T) {
		assert.Equal(t, true, engine.NewPoint(0, 0).IsValid(10))
		assert.Equal(t, true, engine.NewPoint(5, 5).IsValid(10))
		assert.Equal(t, true, engine.NewPoint(10, 10).IsValid(10))
		assert.Equal(t, true, engine.NewPoint(0, 10).IsValid(10))
		assert.Equal(t, true, engine.NewPoint(10, 10).IsValid(10))
	})
	t.Run("invalid points", func(t *testing.T) {
		assert.Equal(t, false, engine.NewPoint(-1, 10).IsValid(10))
		assert.Equal(t, false, engine.NewPoint(10, -1).IsValid(10))
		assert.Equal(t, false, engine.NewPoint(11, 10).IsValid(10))
		assert.Equal(t, false, engine.NewPoint(10, 11).IsValid(10))
	})
}
