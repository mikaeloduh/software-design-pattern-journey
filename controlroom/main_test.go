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
		c.BindCommand("q", MoveForwardTankCommand{tank: Tank{Writer: &writer}})
		c.BindCommand("w", ConnectTelecomCommand{telecom: Telecom{Writer: &writer}})

		c.Input("q")

		assert.Equal(t, "The tank has moved forward.\n", writer.String())
		writer.Reset()

		c.Input("w")
		assert.Equal(t, "The telecom has been turned on.\n", writer.String())
	})

	t.Run("test Undo", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		c.BindCommand("q", MoveForwardTankCommand{tank: Tank{Writer: &writer}})

		c.Input("q")
		assert.Equal(t, "The tank has moved forward.\n", writer.String())

		writer.Reset()
		c.Undo()
		assert.Equal(t, "The tank has moved backward.\n", writer.String())
	})

	t.Run("test macro and its undo", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewMainController()
		tank := Tank{Writer: &writer}
		telecom := Telecom{Writer: &writer}
		c.BindCommand("k", MoveForwardTankCommand{tank: tank})
		c.BindCommand("j", MoveBackwardCommand{tank: tank})
		c.BindCommand("i", ConnectTelecomCommand{telecom: telecom})
		c.BindCommand("o", DisconnectTelecomCommand{telecom: telecom})
		c.setMacro("q", Macro{
			ConnectTelecomCommand{telecom: telecom},
			MoveForwardTankCommand{tank: tank},
			MoveForwardTankCommand{tank: tank},
			MoveForwardTankCommand{tank: tank},
		})

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
	})
}
