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

	t.Run("test Telecom Connect", func(t *testing.T) {
		var writer bytes.Buffer

		telecom := &Telecom{Writer: &writer}
		telecom.Connect()

		assert.Equal(t, "The telecom has been turned on.\n", writer.String())
	})

	t.Run("test Telecom Disconnect", func(t *testing.T) {
		var writer bytes.Buffer

		telecom := &Telecom{Writer: &writer}
		telecom.Disconnect()

		assert.Equal(t, "The telecom has been turned off.\n", writer.String())
	})

	t.Run("test MainController", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		c.SetCommand("q", MoveForwardTankCommand{tank: Tank{Writer: &writer}})

		c.Input("q")

		assert.Equal(t, "The tank has moved forward.\n", writer.String())
	})

	t.Run("test Undo", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		c.SetCommand("q", MoveForwardTankCommand{tank: Tank{Writer: &writer}})

		c.Input("q")
		assert.Equal(t, "The tank has moved forward.\n", writer.String())

		writer.Reset()
		c.Undo()
		assert.Equal(t, "The tank has moved backward.\n", writer.String())
	})
}
