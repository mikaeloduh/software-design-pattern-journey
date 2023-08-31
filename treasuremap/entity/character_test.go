package entity

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestCharacter_MoveStep(t *testing.T) {
	t.Run("move up", func(t *testing.T) {

		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveStep(Up)

		assert.Equal(t, 6, c.Position.y)
		assert.Equal(t, 5, c.Position.x)
		assert.Same(t, c, g.WorldMap[6][5].character)
	})
}

func FakeNewCharacter(writer io.Writer) *Character {
	var c *Character
	c = &Character{
		Writer: writer,
		MaxHp:  300,
		Hp:     300,
		State:  NewNormalState(c),
		Speed:  1,
	}
	return c
}
