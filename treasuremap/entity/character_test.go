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

		assert.Same(t, g.WorldMap[5][5].Object, c)
	})
}

func TestCharacter_MoveStep(t *testing.T) {
	t.Run("move up", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveUp()

		assert.Equal(t, 6, c.Position.Y)
		assert.Equal(t, 5, c.Position.X)
		assert.Same(t, c, g.WorldMap[6][5].Object)
	})

	t.Run("move down", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveDown()

		assert.Equal(t, 4, c.Position.Y)
		assert.Equal(t, 5, c.Position.X)
		assert.Same(t, c, g.WorldMap[4][5].Object)
	})

	t.Run("move left", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveLeft()

		assert.Equal(t, 5, c.Position.Y)
		assert.Equal(t, 4, c.Position.X)
		assert.Same(t, c, g.WorldMap[5][4].Object)
	})

	t.Run("move right", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)

		c.MoveRight()

		assert.Equal(t, 5, c.Position.Y)
		assert.Equal(t, 6, c.Position.X)
		assert.Same(t, c, g.WorldMap[5][6].Object)
	})
}

func TestCharacter_Attack(t *testing.T) {
	t.Run("when character facing left, attack should cleanup all monster in the font", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := NewAdventureGame(c)
		g.AddObject(&Monster{MaxHp: 10, Hp: 10, Speed: 1}, 5, 9, Left)

		c.Attack()

		assert.Empty(t, g.WorldMap[0][5])
	})

}

func FakeNewCharacter(writer io.Writer) *Character {
	var c *Character
	c = &Character{
		Writer:       writer,
		MaxHp:        300,
		Hp:           300,
		AttackDamage: 999,
		Speed:        1,
		State:        NewNormalState(c),
	}
	return c
}
