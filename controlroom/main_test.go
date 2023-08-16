package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControlRoom(t *testing.T) {

	t.Run("test Tank MoveForward", func(t *testing.T) {
		var writer bytes.Buffer

		tank := &Tank{Writer: &writer}
		tank.MoveForward()

		assert.Equal(t, "The tank has moved forward.\n", writer.String())
	})

	t.Run("test Tank MoveBackward", func(t *testing.T) {
		var writer bytes.Buffer

		tank := &Tank{Writer: &writer}
		tank.MoveBackward()

		assert.Equal(t, "The tank has moved backward.\n", writer.String())
	})
}
