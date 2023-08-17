package main

import (
	"bytes"
	"controlroom/entity"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControlRoom(t *testing.T) {

	t.Run("test Tank MoveForward", func(t *testing.T) {
		var writer bytes.Buffer

		tank := &entity.Tank{Writer: &writer}
		tank.MoveForward()

		assert.Equal(t, "The tank has moved forward.\n", writer.String())
	})

	t.Run("test Tank MoveBackward", func(t *testing.T) {
		var writer bytes.Buffer

		tank := &entity.Tank{Writer: &writer}
		tank.MoveBackward()

		assert.Equal(t, "The tank has moved backward.\n", writer.String())
	})

	t.Run("test Telecom Connect", func(t *testing.T) {
		var writer bytes.Buffer

		telecom := &entity.Telecom{Writer: &writer}
		telecom.Connect()

		assert.Equal(t, "The telecom has been turned on.\n", writer.String())
	})

	t.Run("test Telecom Disconnect", func(t *testing.T) {
		var writer bytes.Buffer

		telecom := &entity.Telecom{Writer: &writer}
		telecom.Disconnect()

		assert.Equal(t, "The telecom has been turned off.\n", writer.String())
	})

	t.Run("test MainController bind and run command", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewMainController()
		c.BindCommand("f", entity.MoveForwardTankCommand{Tank: FakeNewTank(&writer)})
		c.BindCommand("c", entity.ConnectTelecomCommand{Telecom: FakeNewTelecom(&writer)})

		c.Input("f")

		assert.Equal(t, "The tank has moved forward.\n", writer.String())
		writer.Reset()

		c.Input("c")
		assert.Equal(t, "The telecom has been turned on.\n", writer.String())
	})

	t.Run("test Undo and Redo", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewMainController()
		c.BindCommand("f", entity.MoveForwardTankCommand{Tank: FakeNewTank(&writer)})

		c.Input("f")
		assert.Equal(t, "The tank has moved forward.\n", writer.String())

		writer.Reset()
		c.Undo()
		assert.Equal(t, "The tank has moved backward.\n", writer.String())

		writer.Reset()
		c.Redo()
		assert.Equal(t, "The tank has moved forward.\n", writer.String())
	})

	t.Run("test macro and its undo and redo", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewMainController()
		tank := FakeNewTank(&writer)
		telecom := FakeNewTelecom(&writer)
		c.BindCommand("f", entity.MoveForwardTankCommand{Tank: tank})
		c.BindCommand("r", entity.MoveBackwardCommand{Tank: tank})
		c.BindCommand("i", entity.ConnectTelecomCommand{Telecom: telecom})
		c.BindCommand("d", entity.DisconnectTelecomCommand{Telecom: telecom})
		c.BindCommand("q", entity.Macro{
			entity.ConnectTelecomCommand{Telecom: telecom},
			entity.MoveForwardTankCommand{Tank: tank},
			entity.MoveForwardTankCommand{Tank: tank},
			entity.MoveForwardTankCommand{Tank: tank},
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

func FakeNewTank(w io.Writer) *entity.Tank {
	return &entity.Tank{Writer: w}
}

func FakeNewTelecom(w io.Writer) *entity.Telecom {
	return &entity.Telecom{Writer: w}
}
