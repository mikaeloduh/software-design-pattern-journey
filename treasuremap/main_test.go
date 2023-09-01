package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"treasuremap/entity"
)

func TestCharacterStatus(t *testing.T) {

	t.Run("test when character poisoned, HP should -15 for 3 rounds, then return to NormalState", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewPoisonedState(c))

		assert.Equal(t, 300, c.Hp)
		assert.IsType(t, &entity.PoisonedState{}, c.State)

		g.StartRound()
		assert.Equal(t, 300-15, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-15-15, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-15-15-15, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-15-15-15, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test character in InvincibleState should invulnerable from damage", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewInvincibleState(c))
		c.TakeDamage(10)

		assert.IsType(t, &entity.InvincibleState{}, c.State)
		assert.Equal(t, 300, c.Hp)

		g.StartRound()
		c.TakeDamage(10)

		assert.Equal(t, 300, c.Hp)
		assert.IsType(t, &entity.InvincibleState{}, c.State)

		g.StartRound()
		c.TakeDamage(10)

		assert.Equal(t, 300-10, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test character in HealingState should restore 30 Hp for 5 rounds", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.TakeDamage(200)
		c.SetState(entity.NewHealingState(c))

		g.StartRound()
		assert.Equal(t, 300-200+30, c.Hp)
		assert.IsType(t, &entity.HealingState{}, c.State)

		g.StartRound()
		assert.Equal(t, 300-200+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30+30+30, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test while character's Hp is fully restored in HealingState, then should recover to NormalState immediately", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.TakeDamage(10)

		c.SetState(entity.NewHealingState(c))
		g.StartRound()

		assert.Equal(t, 300, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test while in AcceleratedState, character can take 2 actions per round", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewAcceleratedState(c))
		g.StartRound()

		assert.IsType(t, &entity.AcceleratedState{}, c.State)
		assert.Equal(t, "take action\ntake action\n", writer.String())
	})

	t.Run("test AccelerateState will be interrupted by attack", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewAcceleratedState(c))

		assert.IsType(t, &entity.AcceleratedState{}, c.State)
		assert.Equal(t, 2, c.Speed)

		g.StartRound()
		c.TakeDamage(1)

		assert.IsType(t, &entity.NormalState{}, c.State)
		assert.Equal(t, 1, c.Speed)
	})

	t.Run("test during OrderlessState character can only move two directions", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		_ = entity.NewAdventureGame(c)

		c.SetState(entity.NewOrderlessState(c))

		assert.IsType(t, &entity.OrderlessState{}, c.State)
	})

	t.Run("test StockpileState", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		_ = entity.NewAdventureGame(c)

		c.SetState(entity.NewStockpileState(c))

		assert.IsType(t, &entity.StockpileState{}, c.State)
	})

	t.Run("test StockpileState enter EruptingState after 3 rounds", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewStockpileState(c))

		g.StartRound()
		g.StartRound()
		g.StartRound()

		assert.IsType(t, &entity.EruptingState{}, c.State)
	})

	t.Run("while in StockpileState, characters' attacks are game-wide", func(t *testing.T) {
		var writer bytes.Buffer

		c := FakeNewCharacter(&writer)
		g := entity.NewAdventureGame(c)
		g.AddObject(&entity.Monster{MaxHp: 10, Hp: 10, Speed: 1}, 5, 9, entity.Left)
		g.AddObject(&entity.Monster{MaxHp: 10, Hp: 10, Speed: 1}, 0, 0, entity.Down)

		c.SetState(entity.NewEruptingState(c))

		c.Attack()

		assert.Empty(t, g.WorldMap[9][5])
		assert.Empty(t, g.WorldMap[0][0])
	})

}

func FakeNewCharacter(writer io.Writer) *entity.Character {
	var c *entity.Character
	c = &entity.Character{
		Writer: writer,
		MaxHp:  300,
		Hp:     300,
		State:  entity.NewNormalState(c),
		Speed:  1,
	}
	return c
}
