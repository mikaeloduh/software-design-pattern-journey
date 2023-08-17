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

	t.Run("test MainController bind and run command", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		c.BindCommand("f", MoveForwardTankCommand{tank: Tank{Writer: &writer}})
		c.BindCommand("c", ConnectTelecomCommand{telecom: Telecom{Writer: &writer}})

		c.Input("f")

		assert.Equal(t, "The tank has moved forward.\n", writer.String())
		writer.Reset()

		c.Input("c")
		assert.Equal(t, "The telecom has been turned on.\n", writer.String())
	})

	t.Run("test Undo", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		c.BindCommand("f", MoveForwardTankCommand{tank: Tank{Writer: &writer}})

		c.Input("f")
		assert.Equal(t, "The tank has moved forward.\n", writer.String())

		writer.Reset()
		c.Undo()
		assert.Equal(t, "The tank has moved backward.\n", writer.String())
	})

	t.Run("test macro and its undo and redo", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		tank := Tank{Writer: &writer}
		telecom := Telecom{Writer: &writer}
		c.BindCommand("f", MoveForwardTankCommand{tank: tank})
		c.BindCommand("r", MoveBackwardCommand{tank: tank})
		c.BindCommand("i", ConnectTelecomCommand{telecom: telecom})
		c.BindCommand("d", DisconnectTelecomCommand{telecom: telecom})
		c.BindCommand("q", Macro{
			ConnectTelecomCommand{telecom: telecom},
			MoveForwardTankCommand{tank: tank},
			MoveForwardTankCommand{tank: tank},
			MoveForwardTankCommand{tank: tank},
		})

		c.Input("f")
		writer.Reset()

		c.Input("q")

		assert.Equal(t, "The telecom has been turned on.\n"+
			"The tank has moved forward.\n"+
			"The tank has moved forward.\n"+
			"The tank has moved forward.\n", writer.String())
		writer.Reset()

		c.Undo()
		assert.Equal(t, "The tank has moved backward.\n"+
			"The tank has moved backward.\n"+
			"The tank has moved backward.\n"+
			"The telecom has been turned off.\n", writer.String())
		writer.Reset()

		c.Redo()
		assert.Equal(t, "The telecom has been turned on.\n"+
			"The tank has moved forward.\n"+
			"The tank has moved forward.\n"+
			"The tank has moved forward.\n", writer.String())
	})
}
