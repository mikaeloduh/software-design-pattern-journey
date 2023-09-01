package entity

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestCharacter_SetPosition(t *testing.T) {
	t.Run("position", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		assert.Same(t, g.WorldMap[5][5].object, c)
	})
}

func TestCharacter_MoveStep(t *testing.T) {
	t.Run("move up", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveUp()

		assert.Equal(t, 6, c.Position.y)
		assert.Equal(t, 5, c.Position.x)
		assert.Same(t, c, g.WorldMap[6][5].object)
	})

	t.Run("move down", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveDown()

		assert.Equal(t, 4, c.Position.y)
		assert.Equal(t, 5, c.Position.x)
		assert.Same(t, c, g.WorldMap[4][5].object)
	})

	t.Run("move left", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveLeft()

		assert.Equal(t, 5, c.Position.y)
		assert.Equal(t, 4, c.Position.x)
		assert.Same(t, c, g.WorldMap[5][4].object)
	})

	t.Run("move right", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveRight()

		assert.Equal(t, 5, c.Position.y)
		assert.Equal(t, 6, c.Position.x)
		assert.Same(t, c, g.WorldMap[5][6].object)
	})
}

func TestCharacter_Attack(t *testing.T) {
	t.Skip()
	t.Run("when object facing left, attack should cleanup all monster in the font", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		//c.Attack()

		assert.Empty(t, g.WorldMap[5][5:])
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
